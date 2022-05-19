package main

var (

	//app info
	versionInfo      = "0.0.10"
	adminContact     = "none"
	siteURL          = "https://lxmxnxl.com"
	siteTitle        = "lmnl."
	siteDescription  = "conversations in liminal space"
	siteMediaPreview = "https://user-images.githubusercontent.com/96031819/168491238-5141d096-1dcd-41fb-bcd5-29bab702e8bc.jpg"

	// this is the max number of posts in the entire system
	// it includes tags and self posts
	maxGlobalPosts = 1000

	// this is the max number of posts on the front page
	maxPostsOnMainFeed = 100

	// this is the max number of posts a user can see when browsing self posts
	maxPostsOnSelfFeed = 100

	// this is the max number of tag topic communities at any given time
	maxTags = 100

	// this is the phrase used to authenticate an admin
	adminTicket = ""
)

//
