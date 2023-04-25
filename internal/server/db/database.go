package db

import (
	"github.com/glebarez/sqlite"
	"gitlab.com/texm/shokku/internal/models"
	"gorm.io/gorm"
)

func Init(dsn string) (*gorm.DB, error) {
	dbCfg := &gorm.Config{
		Logger: Logger{},
	}

	/*if cfg.DebugMode == false {
		dbCfg.Logger = dbCfg.Logger.LogMode(logger.Silent)
	}*/

	db, err := gorm.Open(sqlite.Open(dsn), dbCfg)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Server{},
		&models.ServerSecrets{},
		&models.App{},
		&models.Service{},
		&models.User{},
		&models.SSHKey{},
		&models.GithubApp{},
		&models.AppSetupConfig{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
