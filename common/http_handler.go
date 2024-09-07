package common

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/silaselisha/breathe/pkg/types"
)

func HTTPRequest(method, url string,
	payload io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(
		method,
		url,
		payload,
	)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"Content-Type",
		"application/x-www-form-urlencoded",
	)
	return req, nil
}

func HTTPResponse[T *types.ATokenParam | *types.Tracks](req *http.Request) (*T, error) {
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Panic(err)
			return
		}
	}()

	resp_byte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	payload := new(T)
	if err := json.Unmarshal(resp_byte, payload); err != nil {
		return nil, err
	}

	return payload, nil
}
