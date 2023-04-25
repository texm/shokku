package server

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"golang.org/x/crypto/ssh"
	"gorm.io/gorm"

	"gitlab.com/texm/shokku/internal/models"
	"gitlab.com/texm/shokku/internal/server/db"
)

func Bootstrap() error {
	cfg, cfgErr := LoadConfig()
	if cfgErr != nil {
		return fmt.Errorf("failed to load server config: %w", cfgErr)
	}

	var s models.ServerSecrets
	s.SigningKey = []byte(generateRandomString(32))

	svDb, dbErr := db.Init(cfg.DBPath)
	if dbErr != nil {
		return fmt.Errorf("failed to init db: %w", dbErr)
	}

	deleteErr := svDb.Unscoped().
		Session(&gorm.Session{AllowGlobalUpdate: true}).
		Delete(&models.ServerSecrets{}).
		Error
	if deleteErr != nil {
		return fmt.Errorf("failed to delete existing keys: %w", deleteErr)
	}

	key, genErr := rsa.GenerateKey(rand.Reader, 4096)
	if genErr != nil {
		return fmt.Errorf("failed to generate private key: %w", genErr)
	}

	var buf bytes.Buffer
	if encodeErr := gob.NewEncoder(&buf).Encode(key); encodeErr != nil {
		return fmt.Errorf("failed to encode priv key: %w", encodeErr)
	}
	s.DokkuSSHKeyGob = buf.Bytes()

	publicRsaKey, err := ssh.NewPublicKey(&key.PublicKey)
	if err != nil {
		return err
	}

	if saveErr := svDb.Save(&s).Error; saveErr != nil {
		return fmt.Errorf("failed to save private key: %w", saveErr)
	}

	fmt.Printf("%s", bytes.TrimSpace(ssh.MarshalAuthorizedKey(publicRsaKey)))
	return nil
}
