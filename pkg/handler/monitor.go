package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/silaselisha/breathe/common"
	"github.com/silaselisha/breathe/pkg/types"
)

func CheckPlaylist(token *types.ATokenParam) ([]byte, error) {
	req, err := common.HTTPRequest(
		http.MethodGet,
		os.Getenv("PLAYLIST_URL"),
		nil,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	tracks, err := common.HTTPResponse[*types.Tracks](req)
	if err != nil {
		return nil, err
	}

	payload, err := json.MarshalIndent(tracks, "", " ")
	if err != nil {
		return nil, err
	}

	return payload, nil
}
