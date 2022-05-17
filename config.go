package main

var (

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
