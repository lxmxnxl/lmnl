package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func api() {

	corsAllowedHeaders := []string{
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
		"Access-Control-Allow-Origin",
		"Cache-Control",
		"Content-Security-Policy",
		"Feature-Policy",
		"Referrer-Policy",
		"X-Requested-With",
		"Content-Type"}

	corsOrigins := []string{
		"http://lxmxnxl.com",
		"https://lxmxnxl.com",
		"lxmxnxl.com",
		"http://www.lxmxnxl.com",
		"https://www.lxmxnxl.com",
		"www.lxmxnxl.com",
		"localhost",
		"127.0.0.1",
	}

	corsMethods := []string{
		"GET",
		"POST"}

	headersCORS := handlers.AllowedHeaders(corsAllowedHeaders)

	originsCORS := handlers.AllowedOrigins(corsOrigins)

	methodsCORS := handlers.AllowedMethods(corsMethods)

	router := mux.NewRouter()

	// home handler
	router.HandleFunc("/", homeHandler).Methods("GET")

	// self handler
	router.HandleFunc("/self", selfHandler).Methods("GET")
	router.HandleFunc("/self/", selfHandler).Methods("GET")

	// hashtag handler
	router.HandleFunc("/t", homeHandler).Methods("GET")
	router.HandleFunc("/t/", homeHandler).Methods("GET")

	router.HandleFunc("/t/{thisHashtag}", hashtagHandler).Methods("GET")
	router.HandleFunc("/t/{thisHashtag}/", hashtagHandler).Methods("GET")

	// new post handler
	router.HandleFunc("/post", newPostHandler).Methods("POST")
	router.HandleFunc("/post/", newPostHandler).Methods("POST")

	// public api handler
	router.HandleFunc("/api/posts/all", postFirehoseHandler).Methods("GET")
	router.HandleFunc("/api/posts/all/", postFirehoseHandler).Methods("GET")

	router.HandleFunc("/api/tags/all", tagFirehoseHandler).Methods("GET")
	router.HandleFunc("/api/tags/all/", tagFirehoseHandler).Methods("GET")

	router.HandleFunc("/api/info", infoHandler).Methods("GET")
	router.HandleFunc("/api/info/", infoHandler).Methods("GET")

	// http://localhost:8666/api/admin/delete/jolly-violet-warm-laugh/96de70b40f88c88f030078b8
	router.HandleFunc("/"+adminTicket+"/delete/{itemlabel}", adminHandler).Methods("GET")
	router.HandleFunc("/"+adminTicket+"/delete/{itemlabel}/", adminHandler).Methods("GET")

	// send it brah
	log.Fatal(http.ListenAndServe(":8666", handlers.CORS(headersCORS, originsCORS, methodsCORS)(router)))

}

func adminHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Admin ticket: ", adminTicket)

	// get the var from the url
	vars := mux.Vars(r)

	// get a blank set of posts
	var newPosts []Post

	// get a blank set of tags
	var newTags []string

	// this is the item label we want
	itemLabel := vars["itemlabel"]

	var errorMsg []string

	// loop through the posts we see
	for i, s := range Posts {

		// compare the label of the post we want to remove
		if strings.Compare(itemLabel, s.Label) == 0 {

			// mourn the dead
			log.Println("Item ", i, " marked for deletion.")

			// make a new list but without this one guy
			newPosts = append(Posts[:i], Posts[i+1:]...)

			// Length of Posts
			log.Println("Length of Posts: ", len(Posts))

			// Length of NewPosts
			log.Println("Length of Posts: ", len(newPosts))

			// just for shits and giggles lets check if posts and newposts have the correct number
			if len(newPosts) != len(Posts)-1 {

				errorMsg = append(errorMsg, "Deleted Posts count mismatch. ")

				// for all we know someone could have posted in this time.
				log.Println("Huh, thats weird, did we delete too many?")

			}

			Posts = newPosts

		}

	}

	// loop through the tags we see
	for i, s := range Tags {

		// compare the label of the post we want to remove
		if strings.Compare(itemLabel, s) == 0 {

			// mourn the dead
			log.Println("Item ", i, " marked for deletion.")

			// make a new list but without this one guy
			newTags = append(Tags[:i], Tags[i+1:]...)

			// Length of Posts
			log.Println("Length of Tags: ", len(Tags))

			// Length of newTags
			log.Println("Length of newTags: ", len(newTags))

			// just for shits and giggles lets check if posts and newposts have the correct number
			if len(newTags) != len(Tags)-1 {

				errorMsg = append(errorMsg, "Delete Tags count mismatch. ")

				// for all we know someone could have posted in this time.
				log.Println("Huh, thats weird, did we delete too many?")

			}

			Tags = newTags

		}

	}

	completeErr := strings.Join(errorMsg, "\n")

	fmt.Fprint(w, "\""+completeErr+"\"")

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	// make a little bento for the user and session data to go in neatly
	type homeData struct {
		User  User
		Posts []Post
		Tags  []string
	}

	// we call it hd
	var hd homeData

	// sort out the user session
	hd.User = GetSession(w, r)

	// load up global posts so that we can display them
	hd.Posts = Posts

	// do the same for tags
	hd.Tags, Tags = Unique(Tags), Unique(Tags)

	// gather the files that make up the view
	files := []string{

		"./html/home.html",
	}

	// parse them
	t, parseHomeFiles := template.ParseFiles(files...)

	// look for errs, or in this case Errrs
	Errr(parseHomeFiles)

	// if there are no errs, report the error
	if parseHomeFiles != nil {

		// if for some reason we couldnt make this view happen, shit the bed
		http.Error(w, "Internal Server Error. \n\nmy guy... \ncall the devs, shit is on fire.", http.StatusInternalServerError)

		return

	}

	// a blank set of posts
	var mainPosts []Post

	// look through the global posts, take the ones with tags
	for _, s := range Posts {

		// if there are more than 0 tags...
		if len(s.Tags) > 0 {

			// ..append them
			mainPosts = append(mainPosts, s)

		}

	}

	// out of the ones with tags, make sure we have no more than the limit
	if len(mainPosts) > maxPostsOnMainFeed {

		// when we have more posts than the limit, we take the range of the number of posts minus the limit, then everything after
		// this way we grab the newest max amount of posts and discard the oldest, which are at the top.
		mainPosts = mainPosts[len(mainPosts)-maxPostsOnMainFeed:]

	}

	// only main page a few of them.
	hd.Posts = mainPosts

	// take the data payload and apply it to the theme
	t.Execute(w, hd)

}

func postFirehoseHandler(w http.ResponseWriter, r *http.Request) {

	// heads up, this is gonna be json
	w.Header().Set("Content-Type", "application/json")

	// send it
	json.NewEncoder(w).Encode(Posts)

}

func tagFirehoseHandler(w http.ResponseWriter, r *http.Request) {

	// heads up, this is gonna be json
	w.Header().Set("Content-Type", "application/json")

	// send it
	json.NewEncoder(w).Encode(Tags)

}

func infoHandler(w http.ResponseWriter, r *http.Request) {

	type infoPack struct {
		CurrentGlobalUsersCount int `json:"current_global_users_count,omitempty"`
		CurrentGlobalPostsCount int `json:"current_global_posts_count,omitempty"`
		CurrentGlobalTagsCount  int `json:"current_global_tags_count,omitempty"`
		MaxGlobalPosts          int `json:"max_global_posts,omitempty"`
		MaxPostsOnMainFeed      int `json:"max_posts_on_main_feed,omitempty"`
		MaxPostsOnSelfFeed      int `json:"max_posts_on_self_feed,omitempty"`
		MaxTags                 int `json:"max_tags,omitempty"`
	}

	// make an empty infopack
	var ip infoPack

	ip.CurrentGlobalUsersCount = len(Users)
	ip.CurrentGlobalPostsCount = len(Posts)
	ip.CurrentGlobalTagsCount = len(Tags)
	ip.MaxGlobalPosts = maxGlobalPosts
	ip.MaxPostsOnMainFeed = maxPostsOnMainFeed
	ip.MaxPostsOnSelfFeed = maxPostsOnSelfFeed
	ip.MaxTags = maxTags

	// heads up, this is gonna be json
	w.Header().Set("Content-Type", "application/json")

	// send it
	json.NewEncoder(w).Encode(ip)

}

func selfHandler(w http.ResponseWriter, r *http.Request) {

	// make a little bento for the user and session data to go in neatly
	type homeData struct {
		User  User
		Posts []Post
		Tags  []string
	}

	// make a blank hd payload
	var hd homeData

	// make a blank container for posts
	var thesePosts []Post

	// sort out the user session
	hd.User = GetSession(w, r)

	// for each post we have ....
	for _, s := range Posts {

		// lets find the ones with no tags
		if len(s.Tags) == 0 {

			// then lets append them to our list
			thesePosts = append(thesePosts, s)

		}

	}

	// load up filtered global posts so that we can display them
	hd.Posts = thesePosts

	// nothing wrong with having 1000 posts, but we should only show 100 of them
	if len(hd.Posts) > maxPostsOnSelfFeed {

		// when we have more posts than the limit, we take the range of the number of posts minus the limit, then everything after
		// this way we grab the newest max amount of posts and discard the oldest, which are at the top.
		hd.Posts = hd.Posts[len(hd.Posts)-maxPostsOnSelfFeed:]

	}

	// gather the files that make up the view
	files := []string{

		"./html/home.html",
	}

	// parse them
	t, parseHomeFiles := template.ParseFiles(files...)

	// look for errs, or in this case Errrs
	Errr(parseHomeFiles)

	// if there are no errs, report the error
	if parseHomeFiles != nil {

		// if for some reason we couldnt make this view happen, shit the bed
		http.Error(w, "Internal Server Error. \n\nmy guy... \nthings arent going right in your life.", http.StatusInternalServerError)

		// stop doing things
		return

	}

	t.Execute(w, hd)

}

func hashtagHandler(w http.ResponseWriter, r *http.Request) {

	// read the URL params
	vars := mux.Vars(r)

	// the one after the /t/ is the tag we want
	thisTag := vars["thisHashtag"]

	// make a little bento for the user and session data to go in neatly
	type homeData struct {
		User  User
		Posts []Post
		Tags  []string
	}

	// we call it hd
	var hd homeData

	// blank posts container
	var thesePosts []Post

	// sort out the user session
	hd.User = GetSession(w, r)

	// We need a way to see quickly if we have the desired tag in this post
	for _, s := range Posts {

		// so we take all the posts, and make their tags a single string
		allTagsOneLine := strings.Join(s.Tags, " ")

		// then we see if the tag word is in that string
		if strings.Contains(allTagsOneLine, thisTag) {

			// then we append it
			thesePosts = append(thesePosts, s)

		}

	}

	// load up filtered global posts so that we can display them
	hd.Posts = thesePosts

	// do the same for tags
	hd.Tags = Tags

	// gather the files that make up the view
	files := []string{

		"./html/home.html",
	}

	// parse them
	t, parseHomeFiles := template.ParseFiles(files...)

	// look for errs, or in this case Errrs
	Errr(parseHomeFiles)

	// if there are no errs, report the error
	if parseHomeFiles != nil {

		// things are not okay
		http.Error(w, "Internal Server Error. \n\nmy guy... \nthings arent going right in your life.", http.StatusInternalServerError)

		// stop
		return

	}

	// nothing wrong with having 1000 posts, but we should only show 1000 of them
	var mainPosts []Post

	for _, s := range hd.Posts {
		if len(s.Tags) > 0 {
			mainPosts = append(mainPosts, s)
		}
	}

	if len(mainPosts) > 100 {
		mainPosts = mainPosts[len(mainPosts)-100:]
	}

	// only main page a few of them.
	hd.Posts = mainPosts

	// put it together and show something to the user
	t.Execute(w, hd)

}

func newPostHandler(w http.ResponseWriter, r *http.Request) {

	// read the session from the users cookie
	// if they dont have a session, make one
	thisSession := GetSession(w, r)

	// read the session being reported to us when the form is submitted
	thisPostSession := r.PostFormValue("postSession")

	// just for shits and giggles see if the session reported through the form matches the cookie
	// if strings.Compare(thisSession.Session, thisPostSession) != 0 {

	// i mean, what can we do really? fuck it. report it and move on
	// log.Println("The session in the form was different than the session in the request")

	// }

	// this is the body that is reported through the form
	thisPostBody := r.PostFormValue("postBody")

	// hygiene purposes
	filterBracketL := strings.ReplaceAll(thisPostBody, "<", "&lt;")
	filterBracketR := strings.ReplaceAll(filterBracketL, ">", "&gt;")
	filterBacktick := strings.ReplaceAll(filterBracketR, "`", "&#96;")
	filterBackSlash := strings.ReplaceAll(filterBacktick, "\\", "&#92;")
	filterQuotation := strings.ReplaceAll(filterBackSlash, "\"", "&#x22;")

	// return the clean text
	thisPostBody = filterQuotation

	// lets grab all the posts
	for _, s := range Posts {

		// then we compare the content to this post body
		if strings.Compare(s.Content, thisPostBody) == 0 {

			// and if it is a duplicate of what we already have, stiff arm it
			http.Error(w, "Internal Server Error. \nBruh. That message was a duplicate.", http.StatusInternalServerError)

			// then stop doing things
			return

		}

	}

	// if something is wrong or blank, shit the bed
	if len(thisPostBody) < 1 {

		// I dont know if this is the best way to handle blank requests.
		http.Error(w, "Internal Server Error. \nWhat the hell's going on in here? You think we just waste message bodies like they grow on trees?\nDo better.", http.StatusInternalServerError)

		// stop
		return

		// maybe if the session is blank something is wrong too
	} else if len(thisPostSession) < 1 {

		// idk
		http.Error(w, "Internal Server Error. \nhttps://youtu.be/RfiQYRn7fBg?t=9", http.StatusInternalServerError)
		return

	}

	// parse the form and check for errors
	err := r.ParseForm()

	// err nil kata
	Errr(err)

	// from the post body, get the hashtags
	incomingTags := GetHashtags(thisPostBody)

	// to limit someone sybiling the zeitgest, limit 3
	if len(incomingTags) > 3 {

		// the numbering system for slices is kinda weird
		// i think this is 3 but it could be 4 or 14, who knows
		incomingTags = incomingTags[0:3]

	}

	// append tags to the zeitgeist
	Tags = append(Tags, incomingTags...)

	// If adding these caused us to have more than 100 Tags, trim all but the most recent 100
	if len(Tags) > maxTags {

		// this works in posts so it should work here too.
		// maybe some day make this a function to trim things
		// that probably already exists i bet. :/
		Tags = Tags[len(Tags)-maxTags:]
	}

	// make a post from some of this stuff
	thisPost := GenerateNewPost(thisPostBody, thisSession.Session)

	for _, s := range incomingTags {
		thisPost.Tags = append(thisPost.Tags, string(s))
	}

	// store it into the stream
	Posts = append(Posts, thisPost)

	if len(Posts) > maxGlobalPosts {
		Posts = Posts[len(Posts)-maxGlobalPosts:]
	}

	// send the user to the homepage
	http.Redirect(w, r, "/", http.StatusFound)

}

// GenerateNewPost seemed like a good idea at the time. it takes a body and owner, and makes it into a post object.
func GenerateNewPost(postBody, postOwner string) Post {

	// at the time it seemed clever that this was a one line. it started out much bigger.
	return Post{Content: postBody, Label: GenerateHex(24), Owner: postOwner, Tags: []string{}}

}
