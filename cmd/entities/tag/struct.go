package tag

import "checkaem_server/cmd/entities/post"

type Tag struct {
	Name  string
	Count int
	Posts []*post.Post
}

func New(name string) *Tag {
	return &Tag{
		Name: name,
	}
}

func NewEmpty() *Tag {
	return &Tag{}
}
