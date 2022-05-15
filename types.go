package main

type User struct {
	Label   string `json:"label,omitempty"`
	Session string `json:"session,omitempty"`
}

type Post struct {
	Label   string   `json:"label,omitempty"`
	Owner   string   `json:"owner,omitempty"`
	Content string   `json:"content,omitempty"`
	Tags    []string `json:"tags,omitempty"`
}

var (
	Posts []Post
	Users []User
	Tags  []string
)
