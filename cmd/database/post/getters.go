package post

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/post"
	"context"
)

func Get(username, uuid string) (*post.Post, error) {
	row := database.Connection.QueryRow(context.Background(), getQuery, uuid)

	p := post.NewEmpty()
	err := ScanFullPost(username, row, p)

	if err != nil {
		return nil, err
	}

	tags, err := GetTagNamesByPost(p.Id)

	if err != nil {
		return nil, err
	}

	p.TagNames = tags

	return p, nil
}

func GetByCreator(username string) ([]*post.Post, error) {
	rows, err := database.Connection.Query(context.Background(), getCreatedQuery, username)

	if err != nil {
		return nil, err
	}

	var res []*post.Post

	res, err = ScanFullPosts(username, rows, res)
	res, err = FillTags(res)

	if err != nil {
		return nil, err
	}

	post.Sort(res)

	return res, nil
}

func GetTags(uuid string) ([]string, error) {
	rows, err := database.Connection.Query(context.Background(), getTagNamesQuery, uuid)

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

func GetTagNamesByPost(uuid string) ([]string, error) {

	rows, err := database.Connection.Query(context.Background(), getNamesByPostQuery, uuid)

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

func GetName(uuid string) (string, error) {
	var res string
	err := database.Connection.QueryRow(context.Background(), selectNameQuery, uuid).Scan(&res)
	return res, err
}

func IsBookmarked(username, postId string) (bool, error) {
	rows, err := database.Connection.Query(context.Background(), selectBookmarkQuery, username, postId)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}
