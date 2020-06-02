package bookmark

import (
	"checkaem_server/cmd/database"
	"context"
	"errors"
)

func Insert(username, postId string) error {
	tx, err := database.Connection.Begin(context.Background())

	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), insertQuery, username, postId)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	_, err = tx.Exec(context.Background(), incrementBookmarkCountQuery, postId)

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

func Delete(username, postId string) error {
	tx, err := database.Connection.Begin(context.Background())

	if err != nil {
		return err
	}

	var usernameReturned string
	err = tx.QueryRow(context.Background(), deleteQuery, username, postId).Scan(&usernameReturned)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return err
	}

	if username != usernameReturned {
		_ = tx.Rollback(context.Background())
		return errors.New("usernames aren't equal")
	}

	_, err = tx.Exec(context.Background(), decrementBookmarkCountQuery, postId)

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
