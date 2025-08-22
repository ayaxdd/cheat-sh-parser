package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetResource(s string) ([]byte, error) {
	endpoint := url.QueryEscape(s)
	url := prefix + endpoint

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "curl/7.68.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status)
	}

	body, err := processResponse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}

func processResponse(body io.ReadCloser) ([]byte, error) {
	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	body.Close()

	return RemoveANSI(data), nil
}
