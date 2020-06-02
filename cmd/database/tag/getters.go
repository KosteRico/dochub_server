package tag

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/database/post"
	"checkaem_server/cmd/database/subscription"
	postStruct "checkaem_server/cmd/entities/post"
	"checkaem_server/cmd/entities/tag"
	"context"
)

func Get(username, name string) (*tag.Tag, error) {

	t := tag.NewEmpty()

	err := database.Connection.QueryRow(context.Background(), getQuery, name).
		Scan(&t.Name, &t.Count, &t.SubscribersCount)

	if err != nil {
		return nil, err
	}

	rows, err := database.Connection.Query(context.Background(), getPostsQuery, name)

	if err != nil {
		return nil, err
	}

	var posts []*postStruct.Post

	posts, err = post.ScanFullPosts(username, rows, posts)

	if err != nil {
		return nil, err
	}

	posts, err = post.FillTags(posts)

	postStruct.Sort(posts)

	t.Posts = posts

	subscribed, err := subscription.IsSubscribed(username, name)

	if err != nil {
		return nil, err
	}

	t.Subscribed = subscribed

	return t, nil
}

func GetAll() ([]*tag.Tag, error) {
	rows, err := database.Connection.Query(context.Background(), getAllQuery)

	if err != nil {
		return nil, err
	}

	var tags []*tag.Tag

	for rows.Next() {
		t := tag.NewEmpty()
		err = rows.Scan(&t.Name, &t.Count, &t.SubscribersCount)

		if err != nil {
			return nil, err
		}

		tags = append(tags, t)
	}

	return tags, nil
}
