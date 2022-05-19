package main

import (
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// Errr reduces RSI in go programmers
func Errr(thiserr error) {

	// quoth the gopher, "nevermore..."
	if thiserr != nil {

		log.Println(thiserr)

	}

}

// GenerateHex makes variable length hex strings
func GenerateHex(stringLen int) string {

	// Seed the shit out of this thing.
	rand.Seed(time.Now().UTC().UnixNano())

	// The charset is standard hex.
	const letterBytes = "abcdef0123456789"

	// Make some space
	b := make([]byte, stringLen)

	// Iterate to produce random letter selections
	// from the charset.
	for i := range b {

		// Seed again because it's free.
		rand.Seed(time.Now().UTC().UnixNano())

		// Choose a letter
		b[i] = letterBytes[rand.Intn(len(letterBytes))]

	}

	// Return the ID string we made.
	return string(b)

}

// GetHashtags will take a body of text and extract all hashtags. Returns a slice of tags.
func GetHashtags(str string) []string {

	// this regex should grab the hashtags
	var re = regexp.MustCompile(`(?i)#([a-z0-9]+)`)

	// make a blank array of strings
	var theseResults []string

	// loop through everything we see and run it through the filter
	for _, s := range re.FindAllString(str, -1) {

		// if we find something, strip the pound sign off
		sWithNoPoundSign := strings.TrimPrefix(s, "#")

		// append them to our blank array
		theseResults = append(theseResults, sWithNoPoundSign)

	}

	// make a group of tags
	var tt []string

	// range through the results
	tt = append(tt, theseResults...)

	return tt

}

// Unique removes duplicate entries from a []string
func Unique(s []string) []string {

	// make a map
	check := make(map[string]bool)

	// blank string array
	var result []string

	// checkem
	for _, str := range s {

		if _, ok := check[str]; !ok {

			check[str] = true

			// if it is unique, append it
			result = append(result, str)

		}

	}

	return result

}

func init() {

	// default tags set
	Tags = append(Tags, "hmm", "ask", "why")

	// default posts set
	Posts = append(Posts,
		Post{
			Label:   "BEEFFEA57BA1154A11569710",
			Owner:   "BEEFFEA57BA1154A11569710",
			Content: "this is cool, how much is it going to cost me though #hmm #ask",
			Tags:    []string{"hmm", "ask"},
		},
		Post{
			Label:   "CAFEBABE69420710BA7FEA57",
			Owner:   "CAFEBABE69420710BA7FEA57",
			Content: "the great thing is, nobody knows.. enjoy it while you can #ask #why",
			Tags:    []string{"ask", "why"},
		},
		Post{
			Label:   "CAFEBABE69420710BA7FEA57",
			Owner:   "CAFEBABE69420710BA7FEA57",
			Content: "Once upon a time, I dreamt I was a butterfly, fluttering hither and thither, to all intents and purposes a butterfly.",
			Tags:    []string{},
		})

	// run api in a non blocking goroutine
	go api()

	// announce the admin creds
	announceAdminCreds()

	// do not exit
	select {}
}
