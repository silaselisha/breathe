package main

import (
	"encoding/json"
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

	payload := request_params{
		GrantType:    os.Getenv("GRANT_TYPE"),
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}

	base_url := os.Getenv("BASE_URL")
	values := generate_req_paylod(payload)

	req, err := http_request(http.MethodPost, base_url, strings.NewReader(values.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	access_token, err := http_response[access_token_params](req)
	if err != nil {
		log.Fatal(err)
	}

	playlistData := fetch_breathe_playlist(*access_token)

	var playlist tracks
	err = json.Unmarshal(playlistData, &playlist)
	if err != nil {
		log.Fatalf("Error unmarshaling playlist data: %v", err)
	}

	log.Printf("Fetched playlist with %d tracks", playlist.Tracks.Total)

	// TODO: Implement sanitizer and plagiarism checker
	// sanitizedPlaylist := sanitizePlaylist(playlist)
	// plagiarizedTracks := checkPlagiarism(sanitizedPlaylist)

	redisClient := redis_con()
	defer redisClient.Close()

	// TODO: Implement Redis operations
}

func generate_req_paylod(payload request_params) url.Values {
	values := url.Values{}
	values.Set("client_secret", payload.ClientSecret)
	values.Set("client_id", payload.ClientId)
	values.Set("grant_type", payload.GrantType)
	return values
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"strings"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	payload := request_params{
// 		GrantType:    os.Getenv("GRANT_TYPE"),
// 		ClientId:     os.Getenv("CLIENT_ID"),
// 		ClientSecret: os.Getenv("CLIENT_SECRET"),
// 	}

// 	base_url := os.Getenv("BASE_URL")
// 	values := generate_req_paylod(payload)

// 	req, err := http_request(http.MethodPost, base_url, strings.NewReader(values.Encode()))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	access_token, err := http_response[access_token_params](req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//     // fixed this line
// 	playlist := fetch_breathe_playlist(*access_token)
// 	log.Printf("Fetched playlist with %d tracks", playlist.Tracks.Total)

// 	redisClient := redis_con()
// 	defer redisClient.Close()

// 	// // redisClient for Redis operations
// 	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	// defer cancel()

// 	// // Store playlist in redis illustration
// 	// err = redisClient.Set(ctx, "breathe_playlist", playlist, 0).Err()
// 	// if err != nil {
// 	// log.Printf(" Failed to store playlist in Redis: %v",err )
// 	// } else {
// 	// log.Println("Playlist stored in Redis Succesfully")
// 	// }

// 	// // Fetch Playlist from Redis illustration
// 	// val ,err := redisClient.Get(ctx, "breathe_playlist").Result()
// 	// if err != nil {
// 	// log.Printf("Failed to retreive Playlist from Redis: %v", err)
// 	// } else {
// 	// log.Println("Retreived Playlist from Redis",val)
// 	// }
// }

// func generate_req_paylod(payload request_params) url.Values {
// 	values := url.Values{}
// 	values.Set("client_secret", payload.ClientSecret)
// 	values.Set("client_id", payload.ClientId)
// 	values.Set("grant_type", payload.GrantType)

// 	return values
// }
