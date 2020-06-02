package subscription

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/database/post"
	postObj "checkaem_server/cmd/entities/post"
	"context"
)

func IsSubscribed(username, name string) (bool, error) {
	rows, err := database.Connection.Query(context.Background(), getQuery, username, name)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

func GetPosts(username string) (res []*postObj.Post, err error) {
	rows, err := database.Connection.Query(context.Background(), getPostsQuery, username)

	if err != nil {
		return
	}

	res, err = post.ScanFullPosts(username, rows, res)

	if err != nil {
		return nil, err
	}

	res, err = post.FillTags(res)

	if err != nil {
		return
	}

	postObj.Sort(res)

	return
}

func GetTagNames(username string) (res []string, err error) {
	rows, err := database.Connection.Query(context.Background(), getTagNamesQuery, username)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			return
		}
		res = append(res, s)
	}

	return
}
