package user

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/user"
	"context"
)

func Get(username string) (*user.User, error) {
	u := user.NewEmpty()

	err := database.Connection.QueryRow(context.Background(), selectQuery, username).Scan(&u.Username, &u.Password)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func GetAllNames() ([]string, error) {

	rows, err := database.Connection.Query(context.Background(), selectAllQuery)

	if err != nil {
		return nil, err
	}

	var res []string

	for rows.Next() {
		var str string
		err = rows.Scan(&str)
		if err != nil {
			return nil, err
		}
		res = append(res, str)
	}

	return res, nil
}

func Exists(username string) (bool, error) {
	rows, err := database.Connection.Query(context.Background(), checkIsExistsQuery, username)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
