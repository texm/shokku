package models

import "gorm.io/gorm"

type App struct {
	gorm.Model
	Name        string
	IsSetup     bool
	SetupMethod string
}

type AppSetupConfig struct {
	gorm.Model
	AppID        uint
	DeployBranch string
	RepoURL      string
	RepoGitRef   string
	Image        string
}

type Service struct {
	gorm.Model
	Name                string
	Type                string
	BackupAuthSet       bool
	BackupEncryptionSet bool
	BackupBucket        string
}
