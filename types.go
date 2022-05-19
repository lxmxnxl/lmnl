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

type AppInfo struct {
	VersionInfo      string `json:"version_info"`
	AdminContact     string `json:"admin_contact"`
	SiteURL          string `json:"site_url"`
	SiteTitle        string `json:"site_title"`
	SiteDescription  string `json:"site_description"`
	SiteMediaPreview string `json:"site_media_preview"`
}

type InfoPack struct {
	CurrentGlobalUsersCount int `json:"current_global_users_count,omitempty"`
	CurrentGlobalPostsCount int `json:"current_global_posts_count,omitempty"`
	CurrentGlobalTagsCount  int `json:"current_global_tags_count,omitempty"`
	MaxGlobalPosts          int `json:"max_global_posts,omitempty"`
	MaxPostsOnMainFeed      int `json:"max_posts_on_main_feed,omitempty"`
	MaxPostsOnSelfFeed      int `json:"max_posts_on_self_feed,omitempty"`
	MaxTags                 int `json:"max_tags,omitempty"`
}

type payload struct {
	ThisInfo InfoPack
	ThisApp  AppInfo
}

var (
	Posts []Post
	Users []User
	Tags  []string
)
