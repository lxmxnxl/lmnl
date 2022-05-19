package main

import (
	"net/http"
	"strings"
)

// GetSession checks for a cookie, and if it doesn't find one, it issues one
func GetSession(w http.ResponseWriter, r *http.Request) User {

	// Iterate through the cookies we can see
	for _, cookie := range r.Cookies() {

		// For each cookie we find, we should see if it is something we recognize
		if strings.Compare(cookie.Name, "LIMINAL_SPACE_SESSION") == 0 {

			// The cookie was found, and we recognized its type, so now lets see if it is real
			for _, thisUser := range Users {

				// Compare the cookie we have to the ones we know
				if strings.Compare(thisUser.Session, cookie.Value) == 0 {

					// If it matched a known cookie, then return the whole user
					return thisUser

				}
			}
		}
	}

	// If we make it to this point, we've got a new friend, and it's cookie time...

	// form the cookie
	c := http.Cookie{Name: "LIMINAL_SPACE_SESSION", Value: GenerateHex(24), SameSite: http.SameSiteStrictMode}

	// feed cookie to user
	http.SetCookie(w, &c)

	// make a new user with this cookie
	thisGuy := GenerateNewUserFromCookie(c.Value)

	// add this guy to our user list
	Users = append(Users, thisGuy)

	// we are done here.
	return thisGuy
}

// GenerateNewUserFromCookie is your basic casting of cookie to user for the user creation process
func GenerateNewUserFromCookie(cookieBody string) User {

	// this started out at like 12 lines btw.
	return User{Session: cookieBody, Label: GenerateHex(24)}

}
