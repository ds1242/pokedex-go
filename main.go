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
		},
	}
	// fmt.Println(commands)
	for {
		userInput := userPrompt("Pokedex >")
		if userInput == commands[userInput].name {
			commands[userInput].callback()
		} else {
			fmt.Println("Invalid Command. Please Try Again")
			fmt.Println("To see available options, enter help")
		}
	}
}

func userPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Commands Available:")
	fmt.Println("help: Displays this help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye!")
	os.Exit(0)
	return nil
}