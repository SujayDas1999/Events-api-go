package db

import (
	secrets_reader "events-api/helpers/secrets-reader"
	eventModel "events-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := secrets_reader.SecretsReader("config.json")

	var emptySecretsManager secrets_reader.Secrets

	if dsn == emptySecretsManager {
		panic("Unable to retrieve query string")
		return &gorm.DB{}
	}

	database, err := gorm.Open(postgres.Open(dsn.PostgresConn), &gorm.Config{})

	if err != nil {
		panic(err)
		return &gorm.DB{}
	}

	err = database.AutoMigrate(&eventModel.Event{}, &eventModel.User{}, &eventModel.Registration{})

	if err != nil {
		panic(err)
		return &gorm.DB{}
	}

	DB = database

	return DB
}
