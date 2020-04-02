package subscription

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/database/post"
	postObj "checkaem_server/cmd/entities/post"
)

func GetPosts(username string) (res []*postObj.Post, err error) {
	rows, err := database.Connection.Query(getPostsQuery, username)

	if err != nil {
		return
	}

	res, err = post.ScanFullPosts(rows, res)
	res, err = post.FillTags(res)

	if err != nil {
		return
	}

	return
}

func GetTagNames(username string) (res []string, err error) {
	rows, err := database.Connection.Query(getTagNamesQuery, username)

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
