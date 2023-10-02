package models

import (
	"fmt"
	"go-ws-server/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DbConnect() error {
	dns := fmt.Sprintf("host=%v ", config.Config.DATABASE_HOST)
	dns += fmt.Sprintf("port=%v ", config.Config.DATABASE_PORT)
	dns += fmt.Sprintf("user=%v ", config.Config.DATABASE_USER)
	dns += fmt.Sprintf("password=%v ", config.Config.DATABASE_PASSWORD)
	dns += fmt.Sprintf("dbname=%v ", config.Config.DATABASE_DB_NAME)
	dns += fmt.Sprintf("sslmode=%v ", config.Config.DATABASE_SSL_MODE)
	dns += fmt.Sprintf("TimeZone=%v ", config.Config.DATABASE_TIMEZONE)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dns,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		IgnoreRelationshipsWhenMigrating:         true,
	})
	if err != nil {
		return err
	}

	DB = db

	return nil
}

type findAllUsers struct {
	Users []Users `json:"users"`
	Count int64   `json:"count"`
}

type findAllWebhook struct {
	Webhooks []Webhooks `json:"users"`
	Count    int64      `json:"count"`
}
