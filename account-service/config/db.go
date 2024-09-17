package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PostgresConfiguration struct {
	Url      string
	Username string
	Password string
	Database string
}

func GetEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func ParseConfiguration() PostgresConfiguration {
	POSTGRES_URL := GetEnvVariable("POSTGRES_URL")
	POSTGRES_USERNAME := GetEnvVariable("POSTGRES_USERNAME")
	POSTGRES_PASSWORD := GetEnvVariable("POSTGRES_PASSWORD")
	POSTGRES_DB_NAME := GetEnvVariable("POSTGRES_DB_NAME")
	var nc = PostgresConfiguration{
		Url:      POSTGRES_URL,
		Username: POSTGRES_USERNAME,
		Password: POSTGRES_PASSWORD,
		Database: POSTGRES_DB_NAME,
	}
	return nc
}

// TODO: ACTUALLY CONNECT TO THE DATABASE
// func NewDriver(nc PostgresConfiguration) (neo4j.DriverWithContext, error) {
// 	return neo4j.NewDriverWithContext(nc.Url, neo4j.BasicAuth(nc.Username, nc.Password, ""))
// }
