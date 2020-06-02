package user

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/user"
	"context"
)

func Insert(user *user.User) error {
	tx, err := database.Connection.Begin(context.Background())

	if err != nil {
		return err
	}
	defer tx.Commit(context.Background())

	_, err = tx.Exec(context.Background(), insertQuery, user.Username, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func Delete(username string) error {
	tx, err := database.Connection.Begin(context.Background())

	if err != nil {
		return err
	}

	defer tx.Commit(context.Background())

	_, err = tx.Exec(context.Background(), "delete from \"user\" where username = $1", username)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	return nil
}
