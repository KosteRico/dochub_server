package user

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/user"
)

func Insert(user *user.User) error {
	tx, err := database.Conn.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into user values ($1, $2)", user.Username, user.Password)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func Delete(username string) error {
	tx, err := database.Conn.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec("delete from user where username = $1", username)

	if err != nil {
		return err
	}

	return tx.Commit()
}
