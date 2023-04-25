package env

import (
	"github.com/glebarez/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/texm/dokku-go"
	"gitlab.com/texm/shokku/internal/server/dto"
	"gorm.io/gorm"
)

func NewTestingEnvironment() *Env {
	router := echo.New()
	router.Validator = dto.NewRequestValidator()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	dokkuClient := &mockDokkuClient{}

	return &Env{
		Router:    router,
		DB:        db,
		DebugMode: true,
		Dokku:     &dokkuClient.SSHClient,
	}
}

type mockDokkuClient struct {
	dokku.SSHClient

	returnVal string
}

func (mc *mockDokkuClient) Exec(cmd string) (string, error) {
	return cmd, nil
}

func (mc *mockDokkuClient) SetReturnValue(val string) {
	mc.returnVal = val
}
