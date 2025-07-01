package main

import (
	"fmt"
	"io"
	"net/http"
)

type config struct {
	Next     string
	Previous string
}

type LocationAreaAPIResource struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
}

func getBody(url string) ([]byte, error) {
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	if req.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s", req.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}
