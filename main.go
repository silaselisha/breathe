package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"

	"github.com/silaselisha/breathe/common"
	"github.com/silaselisha/breathe/pkg/handler"
	"github.com/silaselisha/breathe/pkg/types"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	payload := types.ReqParam{
		GrantType:    os.Getenv("GRANT_TYPE"),
		ClientId:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
	}

	base_url := os.Getenv("BASE_URL")
	values := reqPayloadGen(&payload)

	req, err := common.HTTPRequest(
		http.MethodPost,
		base_url,
		strings.NewReader(values.Encode()),
	)
	if err != nil {
		panic(err)
	}

	token, err := common.HTTPResponse[*types.ATokenParam](req)
	if err != nil {
		panic(err)
	}

	_, err = handler.CheckPlaylist(*token)
	if err != nil {
		panic(err)
	}
}

func reqPayloadGen(pld *types.ReqParam) url.Values {
	values := url.Values{}
	values.Set("client_secret", pld.ClientSecret)
	values.Set("client_id", pld.ClientId)
	values.Set("grant_type", pld.GrantType)

	return values
}
