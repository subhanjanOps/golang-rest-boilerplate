package configs

import (
	"errors"
	"log"
	"os"

	md "golang-rest/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"

	// will work on it later TODO: activate logger according to APP_ENV
	// var gormOpts *gorm.Config

	// if os.Getenv("APP_ENV") == "LOCAL" {
	gormOpts := &gorm.Config{
		// Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	// } else {
	// 	gormOpts = &gorm.Config{
	// 		Logger: logger.New(),
	// 	}
	// }

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: dsn,
		}),
		gormOpts,
	)
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	} else {
		log.Println("Connected to db ...")
	}

	return db
}
func AutoMigrate(DB *gorm.DB) error {
	if err := DB.AutoMigrate(md.User{}); err != nil {
		return errors.New("Unable autoMigrateDB - " + err.Error())
	}
	return nil
}

var DB *gorm.DB
