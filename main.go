package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {

	cleanText := strings.Fields(text)
	for i, v := range cleanText {
		cleanText[i] = strings.TrimSpace(v)
		cleanText[i] = strings.ToLower(v)
	}
	return cleanText
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleanText := cleanInput(text)
		if cmd, ok := commands[cleanText[0]]; ok {
			args := cleanText[1:]
			cmd.callback(cfg, args)
		} else {
			fmt.Print("Unknown command \n")
		}
	}
}
