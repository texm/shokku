package models

import (
	"gorm.io/gorm"
)

type GithubApp struct {
	gorm.Model
	AppId         int64
	NodeId        string
	ClientId      string
	Name          string
	Slug          string
	ClientSecret  string
	WebhookSecret string
	PEM           string
}

func (GithubApp) TableName() string {
	return "github_app"
}
