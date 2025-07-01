package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/m-pawlicki/pokedex-cli/internal/pokecache"
)

type Config struct {
	Next     string
	Previous string
	cache    *pokecache.Cache
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

func (cfg *Config) getBody(url string) ([]byte, error) {
	v, ok := cfg.cache.Get(url)
	if ok {
		return v, nil
	}
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	if req.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s", req.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	cfg.cache.Add(url, body)
	return body, nil
}
