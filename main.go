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

	// TODO: Implementing sanitizer and plagiarism checker
	sanitizedPlaylist := sanitizePlaylist(playlist)
	plagiarizedTracks := checkPlagiarism(sanitizedPlaylist)

	linkedList := &LinkedList[item]{}
	for _, item := range sanitizePlaylist.Tracks.Items {
	linkedList.push(item)
	}

	redisClient := redis_con()
	defer redisClient.Close()

	// TODO: Implement Redis operations
	//
	// Storing sanitized Playlists
    sanitizedPlaylistJSON, _ := json.Marshal(sanitizedPlaylist)
    err = redisClient.Set(ctx, "sanitized_playlist", sanitizedPlaylistJSON, 0).Err()
    if err != nil {
        log.Printf("Failed to store sanitized playlist in Redis: %v",err)
    // Storing plagiarized tracks on redis
    plagiarizedTracksJSON, _ := json.Marshal(plagiarizedTracks)
    err = redisClient.Set(ctx, "plagiarized_tracks", plagiarizedTracksJSON, 0).Err()
    if err != nil {
        log.Printf("Failed to store plagiarized tracks in Redis: %v", err)
        }
        log.Printf("Number of plagiarized tracks: %d", len(plagiarizedTracks))
        for _, track := range plagiarizedTracks {
            log.Printf("Plagiarized track: %s", track)
        }
// TODO: Implementing a feedback loop from Redis to breathe component
// // TODO: Implementing a flush redisdb button (maybe)?
}

func generate_req_paylod(payload request_params) url.Values {
	values := url.Values{}
	values.Set("client_secret", payload.ClientSecret)
	values.Set("client_id", payload.ClientId)
	values.Set("grant_type", payload.GrantType)
	return values
}

http.HandleFunc("/playlist", func(w http.ResponseWriter, r *http.Request) {
    sanitizedPlaylistJSON, err := redisClient.Get(ctx, "sanitized_playlist").Result()
    if err != nil {
        http.Error(w, "Failed to retrieve playlist", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(sanitizedPlaylistJSON))
})

http.HandleFunc("/plagiarized", func(w http.ResponseWriter, r *http.Request) {
    plagiarizedTracksJSON, err := redisClient.Get(ctx, "plagiarized_tracks").Result()
    if err != nil {
        http.Error(w, "Failed to retrieve plagiarized tracks", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte(plagiarizedTracksJSON))
})

log.Fatal(http.ListenAndServe(":8080", nil))



















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
