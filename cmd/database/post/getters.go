package post

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/post"
)

func Get(uuid string) (*post.Post, error) {
	row := database.Connection.QueryRow(getQuery, uuid)

	p := post.NewEmpty()
	err := ScanFullPost(row, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func GetByCreator(username string) ([]*post.Post, error) {
	rows, err := database.Connection.Query(getCreatedQuery, username)

	if err != nil {
		return nil, err
	}

	var res []*post.Post

	res, err = ScanFullPosts(rows, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func GetTags(uuid string) ([]string, error) {
	rows, err := database.Connection.Query(getTagNamesQuery, uuid)

	if err != nil {
		return nil, err
	}

	var res []string

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			return nil, err
		}
		res = append(res, s)
	}

	return res, nil
}
