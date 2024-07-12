package main

import (
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

	monitor_breathe_playlist(*access_token)
}

func generate_req_paylod(payload request_params) url.Values {
	values := url.Values{}
	values.Set("client_secret", payload.ClientSecret)
	values.Set("client_id", payload.ClientId)
	values.Set("grant_type", payload.GrantType)

	return values
}
