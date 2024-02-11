package main

import (
	"context"
	"errors"
	"io"
	"net/http"
)

func SendHTTPRequestWithContext(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", errors.New("Something went wrong...")
	}
	client := &http.Client{}
	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return "", errors.New("Something went wrong...")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Something went wrong...")
	}
	return string(body), nil
}
