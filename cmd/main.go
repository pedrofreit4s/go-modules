package main

import (
	"log"

	"github.com/pedrofreit4s/go-modules/config"
	"github.com/pedrofreit4s/go-modules/internal/database"
	"github.com/pedrofreit4s/go-modules/internal/utils"
)

func main() {
	// Load environment variables
	if err := utils.LoadEnvironments(); err != nil {
		log.Fatal("failed to load environment variables: " + err.Error())
	}

	// Initialize configuration
	cnf := config.Init()

	// Initialize database
	database.ConnectToDatabase(cnf)
}
