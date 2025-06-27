package main
import (
	"fmt"
	"os"
	"encoding/json"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *config) error
	config *config
}

var commands = map[string]cliCommand {}
var cfg = &config{}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, val := range commands {
		fmt.Printf("%v : %v \n", val.name, val.description)
	}
	return nil
}

func commandMap(config *config) error {
	if config.Next == "" {
		config.Next = "https://pokeapi.co/api/v2/location-area/"
	}
	body,err := getBody(config.Next)
	if err != nil {
		return err
	}
	res := LocationAreaAPIResource{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	config.Next = res.Next
	config.Previous = res.Previous
	for i := range res.Results {
		fmt.Printf("%s\n", res.Results[i].Name)
	}
	return nil
}

func commandMapBack(config *config) error {
	if config.Previous == "" {
		fmt.Printf("You're on the first page!\n")
		return nil
	}
	body,err := getBody(config.Previous)
	if err != nil {
		return err
	}
	res := LocationAreaAPIResource{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	config.Next = res.Next
	config.Previous = res.Previous
	for i := range res.Results {
		fmt.Printf("%s\n", res.Results[i].Name)
	}
	return nil
}

func init() {

	commands["exit"] = cliCommand{
        name: "exit",
        description: "Exit the Pokedex",
        callback: commandExit,
		config: cfg,
    }

	commands["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
		config: cfg,
	}

	commands["map"] = cliCommand {
		name: "map",
		description: "Displays the next 20 locations",
		callback: commandMap,
		config: cfg,
	}

	commands["mapb"] = cliCommand {
		name: "mapb",
		description: "Displays the previous 20 locations",
		callback: commandMapBack,
		config: cfg,
	}
}