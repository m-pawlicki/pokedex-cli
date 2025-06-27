package main
import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand {}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, val := range commands {
		fmt.Printf("%v : %v \n", val.name, val.description)
	}
	return nil
}

func init() {
	commands["exit"] = cliCommand{
        name: "exit",
        description: "Exit the Pokedex",
        callback: commandExit,
    }

	commands["help"] = cliCommand{
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	}
}