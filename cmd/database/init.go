package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

var Connection *pgxpool.Pool

func Init() error {
	username := os.Getenv("db_username")
	password := os.Getenv("db_password")
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")
	dbName := os.Getenv("db_name")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)

	var err error

	Connection, err = pgxpool.Connect(context.Background(), connStr)

	if err != nil {
		return err
	}

	return nil
}

func Close() {
	Connection.Close()
}
