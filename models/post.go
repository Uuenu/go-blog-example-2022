package models

type Post struct {
	Id       string
	Title    string
	Contenct string
}

func NewPost(id, title, content string) *Post {
	return &Post{id, title, content}
}
