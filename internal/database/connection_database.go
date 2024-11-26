package database

import (
	"fmt"
	"log"

	"github.com/pedrofreit4s/go-modules/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(cnf *config.Config) *gorm.DB {
	// Create the connection DSN
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cnf.DBHost, cnf.DBPort, cnf.DBUser, cnf.DBPassword, cnf.DBName, cnf.SSLMode,
	)

	// Open the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	return db
}
