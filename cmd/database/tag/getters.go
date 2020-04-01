package tag

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/database/post"
	postStruct "checkaem_server/cmd/entities/post"
	"checkaem_server/cmd/entities/tag"
)

func Get(name string) (*tag.Tag, error) {

	t := tag.NewEmpty()

	err := database.Connection.QueryRow(getQuery, name).Scan(&t.Name, &t.Count)

	if err != nil {
		return nil, err
	}

	rows, err := database.Connection.Query(getPostsQuery, name)

	if err != nil {
		return nil, err
	}

	var posts []*postStruct.Post

	posts, err = post.ScanFullPosts(rows, posts)

	if err != nil {
		return nil, err
	}

	t.Posts = posts

	return t, nil
}
