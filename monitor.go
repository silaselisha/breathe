package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func fetch_breathe_playlist(token access_token_params) []byte {
	playlistURL := os.Getenv("PLAYLIST_URL")
	if playlistURL == "" {
		log.Fatal("PLAYLIST_URL is not set in the environment")
	}

	req, err := http.NewRequest(http.MethodGet, playlistURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	return body
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// )

// func fetch_breathe_playlist(token access_token_params) (output []byte) {
//     playlistURL := os.Getenv("PLAYLIST_URL")
//     if playlistURL == "" {
//         log.Fatal("Playlist URL is not set in the environment")
//     }
//    	req, err := http_request(http.MethodGet, playlistURL, nil)
// 	if err != nil {
// 		log.Fatalf("Error creating request: %v", err)
// 	}

// 	req.Header.Set("Authorization",fmt.Sprintf("Bearer "+token.AccessToken))

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 	log.Fatalf("Error making the request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 	log.Fatalf("Unexpected status code: %v",resp.StatusCode)
// 	}

// 	var playlist tracks
// 	err =json.NewDecoder(resp.Body).Decode(&playlist)
// 	if err != nil {
// 	log.Fatalf("Error decoding responses: %v", err)
// 	}

// 	return playlist

// // 	tracks, err := http_response[tracks](req)
// // 	if err != nil {
// // 		log.Panic(err)
// // 		return
// // 	}

// // 	output, err = json.MarshalIndent(tracks, "", " ")
// // 	if err != nil {
// // 		log.Panic(err)
// // 		return
// // 	}
