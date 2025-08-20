package database

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

// InitDB initializes the singleton DB instance
func InitDB() error {
	var err error
	once.Do(func() {
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := os.Getenv("POSTGRES_DB")
		host := os.Getenv("POSTGRES_HOST")
		if host == "" {
			host = "localhost"
		}
		port := os.Getenv("POSTGRES_PORT")
		if port == "" {
			port = "5432"
		}

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	})
	return err
}
