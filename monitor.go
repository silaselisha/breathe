package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func fetch_breathe_playlist(token access_token_params) (output []byte) {
	req, err := http_request(http.MethodGet, os.Getenv("PLAYLIST_URL"), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	tracks, err := http_response[tracks](req)
	if err != nil {
		log.Panic(err)
		return
	}

	output, err = json.MarshalIndent(tracks, "", " ")
	if err != nil {
		log.Panic(err)
		return
	}

	// store tracks in REDIS
	return
}
