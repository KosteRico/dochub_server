package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Connection *pgx.Conn

func Init() error {
	err := godotenv.Load()

	if err != nil {
		log.Println("File \".env\" wasn't initialized")
	}

	username := os.Getenv("db_username")
	password := os.Getenv("db_password")
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")
	dbName := os.Getenv("db_name")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)

	connConfig, err := pgx.ParseConnectionString(connStr)

	if err != nil {
		return err
	}

	Connection, err = pgx.Connect(connConfig)

	if err != nil {
		return err
	}

	return nil
}

func Close() error {
	return Connection.Close()
}
