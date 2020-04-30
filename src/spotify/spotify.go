package spotify

import "net/http"

const baseUrl = "https://api.spotify.com/v1/"

type Client struct {
	http	*http.Client
	baseUrl	string
}

func NewClient(client *http.Client) Client {
	return Client {
		http: 		client,
		baseUrl: 	baseUrl,
	}
}