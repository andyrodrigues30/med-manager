package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func MustGetEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}

	return value
}

func main() {
	// Load environment variables
	envLoadErr := godotenv.Load()
	if envLoadErr != nil {
		log.Fatal("Error loading .env file")
	}

	// Set database connection string
	ConnString := MustGetEnv("DATABASE_URL")

	db := sqlx.MustConnect("postgres", ConnString)

	var version string
	err := db.QueryRow("select version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}
