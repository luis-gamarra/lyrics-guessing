package authenticate

import (
	"golang.org/x/oauth2"
	"os"
)

const (
	AuthUrl = "https://accounts.spotify.com/authorize"
	TokenUrl = "https://accounts.spotify.com/api/token"
)

type Authenticator struct {
	config		*oauth2.Config
}

func NewAuthenticator(redirectUrl string, scopes ...string) Authenticator {
	config := &oauth2.Config {
		ClientID: 		os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: 	os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedirectURL: 	redirectUrl,
		Scopes: 		scopes,
		Endpoint: oauth2.Endpoint {
			AuthURL:  AuthUrl,
			TokenURL: TokenUrl,
		},
	}
	return Authenticator {
		config: config,
	}
}

func (a Authenticator) AuthURL(state string) string {
	return a.config.AuthCodeURL(state)
}