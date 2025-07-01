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
	Cache    *pokecache.Cache
	Pokedex  map[string]Pokemon
}

type LocationAreaAPIResource struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name       string             `json:"name"`
	Encounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name    string        `json:"name"`
	BaseEXP int           `json:"base_experience"`
	Height  int           `json:"height"`
	Weight  int           `json:"weight"`
	Stats   []PokemonStat `json:"stats"`
	Types   []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

type PokemonType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

func (cfg *Config) getBody(url string) ([]byte, error) {
	v, ok := cfg.Cache.Get(url)
	if ok {
		return v, nil
	}
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	if req.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and\nbody: %s", req.StatusCode, body)
	}
	cfg.Cache.Add(url, body)
	return body, nil
}
