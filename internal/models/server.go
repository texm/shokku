package models

import (
	"gitlab.com/texm/shokku/internal/server/auth"
	"gorm.io/gorm"
	"time"
)

type Server struct {
	gorm.Model
	IsSetup    bool
	SetupKey   string
	AuthMethod auth.Method
	LastSync   time.Time
}

func (Server) TableName() string {
	return "server"
}

type ServerSecrets struct {
	gorm.Model
	DokkuSSHKeyGob []byte
	SigningKey     []byte
}
