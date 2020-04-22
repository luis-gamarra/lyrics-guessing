package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AvailableDevices struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	ID       string `json:"id"`
	IsActive bool   `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}

const (
	spotify_token                = "SPOTIFY_TOKEN"
	authToken                    = "Bearer "
	userAvailableDevicesEndpoint = "https://api.spotify.com/v1/me/player/devices"
)

func main() {
	body := makeGetAPICall(userAvailableDevicesEndpoint)

	fmt.Println(string(body))
}

func makeGetAPICall(endpoint string) []byte {

	client := http.Client{}

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		log.Fatal(err)
	}

	authHeader := fmt.Sprintf("%s%s", authToken, os.Getenv(spotify_token))

	request.Header.Set("Authorization", authHeader)

	response, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}
