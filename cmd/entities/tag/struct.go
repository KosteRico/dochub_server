package tag

import "checkaem_server/cmd/entities/post"

type Tag struct {
	Name             string       `json:"name"`
	Count            uint         `json:"count"`
	SubscribersCount uint         `json:"subscribers_count"`
	Subscribed       bool         `json:"subscribed,omitempty"`
	Posts            []*post.Post `json:"posts,omitempty"`
}

func New(name string) *Tag {
	return &Tag{
		Name: name,
	}
}

func NewEmpty() *Tag {
	return &Tag{}
}
