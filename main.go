package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}

	payload := request_params{
		GrantType:    os.Getenv("GRANT_TYPE"),
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}

	values := url.Values{}
	values.Set("client_secret", payload.ClientSecret)
	values.Set("client_id", payload.ClientId)
	values.Set("grant_type", payload.GrantType)

	base_url := os.Getenv("BASE_URL")

	req, err := http.NewRequest(http.MethodPost, base_url, strings.NewReader(values.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	token := new(access_token_params)
	err = json.Unmarshal(res, &token)
	if err != nil {
		log.Fatal(err)
	}

	// manage and monitor BREATHE PLAYLIST
	monitorBreathePlaylist(*token)
}

func monitorBreathePlaylist(token access_token_params) {
	// create a request object
	url := "https://api.spotify.com/v1/playlists/7uMISrpOiJ0qJ53f3bpNYB"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	data := new(tracks)
	err = json.Unmarshal(res, &data)
	if err != nil {
		log.Panic(err)
	}

	output, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("breathe.json", output, 0644); err != nil {
		log.Fatal(err)
	}
}
