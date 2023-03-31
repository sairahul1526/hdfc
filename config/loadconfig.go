package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// LoadConfig - loads all env vars
func LoadConfig() error {
	// load .env file from given path for local, else will be getting from env var
	if !strings.EqualFold(os.Getenv("prod"), "true") {
		configFile := ".test-env"
		if strings.EqualFold(os.Getenv("TESTING"), "true") { // to run testcases
			configFile = "../../../.test-env"
		}
		err := godotenv.Load(configFile)
		if err != nil {
			return err
		}
	}

	// main postgres db
	MainDBConfig = os.Getenv("MAIN_DB_CONFIG")
	MainDBDatabase = os.Getenv("MAIN_DB_DATABASE")
	Log, _ = strconv.ParseBool(os.Getenv("LOG"))
	Migrate, _ = strconv.ParseBool(os.Getenv("MIGRATE"))

	// s3

	// razorpay

	return nil
}
