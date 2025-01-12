package database

import (
	"fmt"
	"gin-gonic-gorm/configs/db_config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	fmt.Println("Connecting to database...")
	var errConnection error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", db_config.DB_HOST, db_config.DB_USER, db_config.DB_PASSWORD, db_config.DB_NAME, db_config.DB_PORT)

	DB, errConnection = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errConnection != nil {
		panic("cannot connect to database")
	}

	log.Println("Connected to database")

}
