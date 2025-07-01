package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/m-pawlicki/pokedex-cli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, args []string) error
	config      *Config
}

var commands = map[string]cliCommand{}
var cfg = &Config{Cache: pokecache.NewCache(time.Minute)}

func commandExit(config *Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, val := range commands {
		fmt.Printf("%v : %v \n", val.name, val.description)
	}
	return nil
}

func commandMap(config *Config, args []string) error {
	if config.Next == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area/"
	}
	body, err := cfg.getBody(config.Next)
	if err != nil {
		fmt.Println(err)
		return err
	}
	res := LocationAreaAPIResource{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
		return err
	}
	config.Next = res.Next
	config.Previous = res.Previous
	for i := range res.Results {
		fmt.Printf("%s\n", res.Results[i].Name)
	}
	return nil
}

func commandMapBack(config *Config, args []string) error {
	if config.Previous == "" {
		fmt.Printf("You're on the first page!\n")
		return nil
	}
	body, err := cfg.getBody(config.Previous)
	if err != nil {
		fmt.Println(err)
		return err
	}
	res := LocationAreaAPIResource{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
		return err
	}
	config.Next = res.Next
	config.Previous = res.Previous
	for i := range res.Results {
		fmt.Printf("%s\n", res.Results[i].Name)
	}
	return nil
}

func commandExplore(config *Config, args []string) error {
	if len(args) < 1 {
		fmt.Printf("Please input an area to explore\nUsage: explore <area-location>\n")
		return nil
	}
	area := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", args[0])
	body, err := cfg.getBody(area)
	if err != nil {
		fmt.Println(err)
		return err
	}
	res := LocationArea{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Exploring %s...\n", args[0])
	if len(res.Encounters) >= 1 {
		fmt.Println("Found Pokemon:")
		for i := range res.Encounters {
			fmt.Printf("- %s\n", res.Encounters[i].Pokemon.Name)
		}
	} else {
		fmt.Println("No Pokemon Found.")
	}
	return nil
}

func commandCatch(config *Config, args []string) error {
	if len(args) < 1 {
		fmt.Printf("Please input a Pokemon to try and catch\nUsage: catch <pokemon-name>\n")
		return nil
	}
	return nil
}

func init() {

	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
		config:      cfg,
	}

	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
		config:      cfg,
	}

	commands["map"] = cliCommand{
		name:        "map",
		description: "Displays the next 20 locations",
		callback:    commandMap,
		config:      cfg,
	}

	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "Displays the previous 20 locations",
		callback:    commandMapBack,
		config:      cfg,
	}

	commands["explore"] = cliCommand{
		name:        "explore <area-name>",
		description: "Explore a specific area for pokemon",
		callback:    commandExplore,
		config:      cfg,
	}

	commands["catch"] = cliCommand{
		name:        "catch <pokemon-name>",
		description: "Attempt to catch a Pokemon",
		callback:    commandCatch,
		config:      cfg,
	}
}
