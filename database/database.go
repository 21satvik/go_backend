package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func connectDb() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to database successfully!")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations...")
	//TODO: Add migrations

	Database = DbInstance{Db: db}
}
