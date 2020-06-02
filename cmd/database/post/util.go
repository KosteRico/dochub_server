package post

import (
	"checkaem_server/cmd/entities/post"
	"github.com/jackc/pgx/v4"
)

func ScanFullPost(usename string, row pgx.Row, p *post.Post) error {

	err := row.Scan(&p.Id, &p.Title, &p.Description, &p.DateCreated,
		&p.DateUpdated, &p.CreatorUsername, &p.BookmarkCount)

	if err != nil {
		return err
	}

	return FillIsBookmarked(usename, p)
}

func ScanFullPosts(username string, rows pgx.Rows, posts []*post.Post) ([]*post.Post, error) {

	defer rows.Close()

	for rows.Next() {
		p := post.NewEmpty()

		err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.DateCreated,
			&p.DateUpdated, &p.CreatorUsername, &p.BookmarkCount)

		if err != nil {
			return nil, err
		}

		err = FillIsBookmarked(username, p)

		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, nil
}

func FillIsBookmarked(username string, p *post.Post) error {
	isBookmarked, err := IsBookmarked(username, p.Id)

	if err != nil {
		return err
	}

	p.IsBookmarked = isBookmarked

	return nil
}

func FillTags(posts []*post.Post) ([]*post.Post, error) {

	var err error

	for _, r := range posts {
		tags, err := GetTagNamesByPost(r.Id)

		if err != nil {
			return nil, err
		}

		r.TagNames = tags
	}

	return posts, err
}
