package database

import (
	CONFIG "hdfc-backend/config"
)

// ConnectDatabases - connects all databases with given configurations
func ConnectDatabases() error {
	// connect to mongo database from config
	err := MongoConnect(CONFIG.MainDBConfig, CONFIG.MainDBDatabase)
	if err != nil {
		return err
	}

	return nil
}
