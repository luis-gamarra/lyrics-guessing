package main

import (
	"fmt"
	"net/http"
	"spotify"
	"spotify/authenticate"
)

const redirectUrl = "http://localhost:8080/callback"

var (
	auth = authenticate.NewAuthenticator(redirectUrl, "user-library-read")
	userChannel = make(chan spotify.Client)
	state = "placeholder"
)

func main() {
	http.HandleFunc("/callback", callbackHandler)

	fmt.Println(auth.AuthURL(state))

	//go http.ListenAndServe(":8080", nil)



}

func callbackHandler(w http.ResponseWriter, r *http.Request) {

}