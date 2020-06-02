package subscription

import (
	"checkaem_server/cmd/database"
	"context"
	"errors"
)

func Insert(username, tagName string) error {
	tx, err := database.Connection.Begin(context.Background())

	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), insertQuery, username, tagName)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	_, err = tx.Exec(context.Background(), incrementCounterQuery, tagName)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	err = tx.Commit(context.Background())

	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	return nil
}

func Delete(username, tagName string) error {
	var usernameReturned string

	tx, err := database.Connection.Begin(context.Background())

	if err != nil {
		return err
	}

	err = tx.QueryRow(context.Background(), deleteQuery, username, tagName).Scan(&usernameReturned)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	if username != usernameReturned {
		return errors.New("usernames aren' equal")
	}

	_, err = tx.Exec(context.Background(), decrementCounterQuery, tagName)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	err = tx.Commit(context.Background())

	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	return nil
}
