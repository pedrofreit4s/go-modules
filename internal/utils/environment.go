package utils

import "github.com/joho/godotenv"

func LoadEnvironments() error {
	// Load all environment variables
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}
