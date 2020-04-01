package subsctiption

import "github.com/jackc/pgx/pgtype"

type Subscription struct {
	Username    string
	TagName     string
	DateCreated pgtype.Timestamptz
}

func New(username, tagName string) *Subscription {
	return &Subscription{
		Username: username,
		TagName:  tagName,
	}
}

func NewEmpty() *Subscription {
	return &Subscription{}
}
