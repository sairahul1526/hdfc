package database

import (
	LOGGER "hdfc-backend/logger"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// mongodb

// MongoConnect - connects to mongodb
// returns error
func MongoConnect(dbConfig, database string) error {
	LOGGER.Log(dbConfig, database)

	err := mgm.SetDefaultConfig(nil, database, options.Client().ApplyURI(dbConfig))
	if err != nil {
		LOGGER.Warn(dbConfig, database, err)
		return err
	}

	LOGGER.Log("Successfully connected to mongodb")
	return nil
}
