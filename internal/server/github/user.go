package github

import (
	"context"
	gh "github.com/google/go-github/v48/github"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/env"
	"gitlab.com/texm/shokku/internal/models"
	"golang.org/x/oauth2"
	ghOAuth "golang.org/x/oauth2/github"
)

type UserClient struct {
	client *gh.Client
}

type CodeExchangeParams struct {
	Code        string
	Scopes      []string
	RedirectURL string
}

func ExchangeCode(ctx context.Context, e *env.Env, p CodeExchangeParams) (*UserClient, error) {
	var ghApp models.GithubApp
	if err := e.DB.WithContext(ctx).First(&ghApp).Error; err != nil {
		log.Error().Err(err).Msg("no github app in db")
		return nil, err
	}

	conf := &oauth2.Config{
		ClientID:     ghApp.ClientId,
		ClientSecret: ghApp.ClientSecret,
		Scopes:       p.Scopes,
		RedirectURL:  p.RedirectURL,
		Endpoint:     ghOAuth.Endpoint,
	}
	token, err := conf.Exchange(ctx, p.Code)
	if err != nil {
		log.Error().Err(err).Msg("failed to exchange code for token")
		return nil, err
	}

	tokenSource := oauth2.StaticTokenSource(token)
	oauthClient := oauth2.NewClient(ctx, tokenSource)
	client := gh.NewClient(oauthClient)
	return &UserClient{client}, nil
}

func (u *UserClient) GetUser(ctx context.Context) (*gh.User, error) {
	user, _, err := u.client.Users.Get(ctx, "")
	if err != nil {
		log.Error().Err(err).Msg("failed to get user")
		return nil, err
	}
	return user, nil
}
