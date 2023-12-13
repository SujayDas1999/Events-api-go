package db

import (
	eventModel "events-api/models/event"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "postgresql://inu________:qrUzdP1RhJ6HdY9aWrT3pQ@test-cockroach-3683.6xw.cockroachlabs.cloud:26257/events-api?sslmode=verify-full"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&eventModel.Event{})

	if err != nil {
		panic(err)
	}

	DB = database

	return database
}

//func (db *DB) SetDB(database *gorm.DB) {
//	db.Db = database
//}
//
//func (db *DB) ReturnCurrentDbInstance() *gorm.DB {
//	return db.Db
//}
