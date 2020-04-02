package post

import (
	"checkaem_server/cmd/entities/post"
	"github.com/jackc/pgx"
)

func ScanFullPost(row *pgx.Row, p *post.Post) error {
	return row.Scan(&p.Id, &p.Title, &p.Description, &p.DateCreated,
		&p.DateUpdated, &p.CreatorUsername, &p.LikeCount, &p.BookmarkCount)
}

func ScanFullPosts(rows *pgx.Rows, posts []*post.Post) ([]*post.Post, error) {

	defer rows.Close()

	for rows.Next() {
		p := post.NewEmpty()
		err := rows.Scan(&p.Id, &p.Title, &p.Description, &p.DateCreated,
			&p.DateUpdated, &p.CreatorUsername, &p.LikeCount, &p.BookmarkCount)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
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
