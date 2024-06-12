package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string
	description string
	callback func() error
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println()
			fmt.Println("Unknown command")
			fmt.Println("enter help to see a list of available commands")
			fmt.Println()
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{		
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map": {
			name:"map",
			description: "Displays the name of 20 locations in the Pokemon world",
			callback: commandMap,
		},
		"mapb": {
			name:"mapb",
			description: "Displays the previous 20 locations queried in map",
			callback: commandMapB,
		},
		
	}
}