package bookmark

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/database/post"
	postObj "checkaem_server/cmd/entities/post"
	"context"
)

func GetPosts(username string) (res []*postObj.Post, err error) {
	rows, err := database.Connection.Query(context.Background(), selectPostsQuery, username)

	if err != nil {
		return nil, err
	}

	res, err = post.ScanFullPosts(username, rows, res)

	if err != nil {
		return nil, err
	}

	res, err = post.FillTags(res)

	if err != nil {
		return nil, err
	}

	return
}
