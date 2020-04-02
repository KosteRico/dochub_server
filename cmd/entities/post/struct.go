package post

import "github.com/jackc/pgx/pgtype"

type Post struct {
	Id              string             `json:"id"`
	Title           string             `json:"title"`
	Description     string             `json:"description"`
	DateCreated     pgtype.Timestamptz `json:"date_created"`
	DateUpdated     pgtype.Timestamptz `json:"date_updated"`
	CreatorUsername string             `json:"creator_username"`
	LikeCount       int                `json:"like_count"`
	BookmarkCount   int                `json:"bookmark_count"`
	TagNames        []string           `json:"tag_names"`
}

func NewEmpty() *Post {
	return &Post{}
}

func New(title, creatorUsername, description string, tags []string) *Post {
	return &Post{
		Title:           title,
		Description:     description,
		CreatorUsername: creatorUsername,
		TagNames:        tags,
	}
}
