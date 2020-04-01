package subscription

import "checkaem_server/cmd/database"

func Insert(username, tagName string) error {
	_, err := database.Connection.Exec(insertQuery, username, tagName)
	return err
}

func Delete(username, tagName string) error {
	_, err := database.Connection.Exec(deleteQuery, username, tagName)
	return err
}
