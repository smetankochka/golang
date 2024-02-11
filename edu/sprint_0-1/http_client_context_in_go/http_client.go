package main

import (
	"errors"
	"io"
	"net/http"
)

func SendHTTPRequest(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", errors.New("Something went wrong...")
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("Something went wrong...")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Something went wrong...")
	}
	return string(body), nil
}
