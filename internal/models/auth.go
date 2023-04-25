package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Source  string
	SSHKeys []SSHKey

	PasswordHash []byte
	TotpEnabled  bool
	TotpSecret   string
	RecoveryCode string
}

type SSHKey struct {
	gorm.Model
	UserID   uint
	GithubID int64 `gorm:"unique"`
	Key      string
}
