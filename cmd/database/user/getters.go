package user

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/user"
)

func Get(username string) (*user.User, error) {
	u := user.NewEmpty()

	err := database.Connection.QueryRow(selectQuery, username).Scan(&u.Username, &u.Password)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func GetAllUsernames() ([]string, error) {

	rows, err := database.Connection.Query(selectAllQuery)

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
