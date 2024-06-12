package main

import "fmt"

type cliCommand struct {
	name string
	description string
	callback func() error
}


func main() {
	commands := map[string]cliCommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		}
	}
	i := fmt.Print("Pokedex > ")
	fmt.Scan(&i)
	fmt.Println(i)
}