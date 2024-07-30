package models

type Post struct {
	id     string `json:"id"`
	title  string `json:"title"`
	userId string `json:"userId"`
}

type PostMetadata struct {
}
