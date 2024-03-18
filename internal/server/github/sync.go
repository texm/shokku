package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	gh "github.com/google/go-github/v48/github"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"gorm.io/gorm/clause"
)

type githubUser struct {
	User    models.User
	SSHKeys []models.SSHKey
}

func Sync(e *env.Env) error {
	if err := SyncUsersToDB(e); err != nil {
		return fmt.Errorf("failed to sync github users: %w", err)
	}

	//if err := SyncInstallationStatus(e); err != nil {
	//	return fmt.Errorf()
	//}
	return nil
}

// SyncUsersToDB asynchronously get users in organization & add to db
func SyncUsersToDB(e *env.Env) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, clientErr := GetAppClient(e)
	if clientErr != nil {
		return fmt.Errorf("failed to get app client: %w", clientErr)
	}

	installs, _, installsErr := client.Apps.ListInstallations(ctx, nil)
	if installsErr != nil {
		return installsErr
	}

	var users []githubUser
	for _, install := range installs {
		var members []*gh.User
		var err error
		var response *gh.Response
		var options *gh.ListMembersOptions
		if install.GetAccount().GetType() == "Organization" {
			var temp_members []*gh.User

			insClient := client.GetInstallationClient(install.GetID())
			org := install.GetAccount().GetLogin()
			temp_members, response, err = insClient.Organizations.ListMembers(ctx, org, options)
			members = append(members, temp_members...)
			
			for response.NextPage != 0 {
				options := &gh.ListMembersOptions{
					ListOptions: gh.ListOptions{
						Page: response.NextPage,
					},
				}
				temp_members, response, err = insClient.Organizations.ListMembers(ctx, org, options)
				members = append(members, temp_members...)
			}

		} else {
			members = append(members, install.GetAccount())
		}
		if err != nil {
			log.Error().Err(err).
				Int64("installation_id", install.GetID()).
				Msg("failed to get members")
			continue
		}
		for _, member := range members {
			users = append(users, fetchUserInfo(member))
		}
	}

	conflict := clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"updated_at"}),
	}

	for _, u := range users {
		if err := e.DB.Clauses(conflict).Create(&u.User).Error; err != nil {
			log.Error().Err(err).
				Str("name", u.User.Name).
				Msg("failed to create user")
			continue
		}
		for _, key := range u.SSHKeys {
			key.UserID = u.User.ID
			if err := e.DB.Clauses(conflict).Create(&key).Error; err != nil {
				log.Error().Err(err).
					Str("name", u.User.Name).
					Msg("failed to create user ssh key")
			}
		}
	}

	oneMinuteAgo := time.Now().Add(-time.Minute)
	var deletedUsers []models.User
	rUsers := e.DB.Where("updated_at < ?", oneMinuteAgo).Delete(&deletedUsers)
	if rUsers.Error != nil {
		log.Error().Err(rUsers.Error).Msg("failed to delete old users")
	}

	var deletedKeys []models.SSHKey
	rKeys := e.DB.Where("updated_at < ?", oneMinuteAgo).Delete(&deletedKeys)
	if rKeys.Error != nil {
		log.Error().Err(rKeys.Error).Msg("failed to delete old ssh keys")
	}

	log.Debug().
		Int("num_installations", len(installs)).
		Int("synced_users", len(users)).
		Int64("removed_users", rUsers.RowsAffected).
		Int64("removed_keys", rKeys.RowsAffected).
		Msgf("github user sync complete")

	return nil
}

func fetchUserInfo(u *gh.User) githubUser {
	username := u.GetLogin()
	user := githubUser{
		User: models.User{Name: username, Source: "github"},
	}
	userKeysApi := fmt.Sprintf("https://api.github.com/users/%s/keys", username)
	res, reqErr := http.Get(userKeysApi)
	if reqErr != nil {
		log.Error().Err(reqErr).
			Str("username", username).
			Msg("failed to get users SSH keys")
		return user
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("failed to read response body")
		return user
	}

	var keys []gh.Key
	if err := json.Unmarshal(body, &keys); err != nil {
		log.Error().Err(err).Msg("failed to unmarshal keys")
		return user
	}

	user.SSHKeys = make([]models.SSHKey, len(keys))
	for i, key := range keys {
		user.SSHKeys[i] = models.SSHKey{
			GithubID: key.GetID(),
			Key:      key.GetKey(),
		}
	}

	return user
}
