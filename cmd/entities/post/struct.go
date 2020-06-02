package post

import (
	"time"
)

type Post struct {
	Id              string    `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	DateCreated     time.Time `json:"date_created"`
	DateUpdated     time.Time `json:"date_updated"`
	CreatorUsername string    `json:"creator_username"`
	IsBookmarked    bool      `json:"is_bookmarked"`
	BookmarkCount   uint      `json:"bookmark_count"`
	TagNames        []string  `json:"tag_names"`
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
