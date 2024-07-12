package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func monitor_breathe_playlist(token access_token_params) {
	req, err := http_request(http.MethodGet, os.Getenv("PLAYLIST_URL"), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	tracks, err := http_response[tracks](req)
	if err != nil {
		log.Panic(err)
	}

	output, err := json.MarshalIndent(tracks, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("breathe.json", output, 0644); err != nil {
		log.Fatal(err)
	}
}
