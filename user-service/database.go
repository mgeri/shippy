package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

const (
	defaultHost     = "localhost"
	defaultUser     = "postgres"
	defaultDBName   = "postgres"
	defaultPassword = "postgres"
)

func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = defaultUser
	}
	DBName := os.Getenv("DB_NAME")
	if DBName == "" {
		DBName = defaultDBName
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = defaultPassword
	}

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, DBName, password,
		),
	)
}
