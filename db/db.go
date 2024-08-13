package db

import (
	"os"

	"github.com/TiveCS/sync-expense/api/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	url, exists := os.LookupEnv("DATABASE_URL")

	if !exists {
		panic("DATABASE_URL is required")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  url,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.User{}, &entities.Account{}, &entities.Transaction{})

	return db
}
