package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/joho/godotenv"
	"os"
)

var Conn *pgx.Conn

func InitDB() error {
	err := godotenv.Load()

	username := os.Getenv("db_username")
	password := os.Getenv("db_password")
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")
	dbName := os.Getenv("db_name")

	if err != nil {
		return err
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)

	connConfig, err := pgx.ParseConnectionString(connStr)

	if err != nil {
		return err
	}

	Conn, err = pgx.Connect(connConfig)

	if err != nil {
		return err
	}

	return nil
}
