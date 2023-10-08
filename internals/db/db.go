package db

import (
	"fmt"
	"log"

	"github.com/developertom01/library-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase() *Database {
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: fmt.Sprintf("password=%v dbname=%v port=%v host=%v user=%v", config.DATABASE_PASSWORD, config.DATABASE_NAME, config.DATABASE_PORT, config.DATABASE_HOST, config.DATABASE_USER)}), &gorm.Config{})
	if err != nil {
		log.Fatal("An error occurred starting db")
	}
	return &Database{
		DB: db,
	}
}
