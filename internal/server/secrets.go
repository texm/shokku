package server

import (
	"bytes"
	"crypto/rsa"
	"database/sql"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"gitlab.com/texm/shokku/internal/models"
	"gorm.io/gorm"
)

type secrets struct {
	signingKey []byte
	privKey    *rsa.PrivateKey
}

func getServerSecrets(db *gorm.DB) (*secrets, error) {
	var s models.ServerSecrets
	if err := db.Find(&s).Error; err != nil && err != sql.ErrNoRows {
		log.Error().
			Err(err).
			Msg("failed to get server secrets")
		return nil, err
	}

	var key *rsa.PrivateKey
	if len(s.DokkuSSHKeyGob) == 0 {
		return nil, errors.New("no ssh key stored")
	}

	r := bytes.NewReader(s.DokkuSSHKeyGob)
	if decodeErr := gob.NewDecoder(r).Decode(&key); decodeErr != nil {
		return nil, fmt.Errorf("failed to decode priv key: %w", decodeErr)
	}

	if validErr := key.Validate(); validErr != nil {
		return nil, fmt.Errorf("private key validation failed: %w", validErr)
	}

	return &secrets{
		privKey:    key,
		signingKey: s.SigningKey,
	}, nil
}
