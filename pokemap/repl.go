package pokemap

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/ds1242/pokedex-go/config"
)

type CliCommand struct {
	name string
	description string
	callback func(*config.Config) error
}

func StartRepl(conf *config.Config) {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Pokedex > ")
		reader.Scan()

		words := CleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := GetCommands(conf)[commandName]
		if exists {
			err := command.callback(conf)
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

func CleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func GetCommands(conf *config.Config) map[string]CliCommand {
	return map[string]CliCommand{		
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: func(conf *config.Config) error { return commandHelp(conf) },
		},
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: func(conf *config.Config) error { return commandExit() },
		},
		"map": {
			name:"map",
			description: "Displays the name of 20 locations in the Pokemon world",
			callback: func(conf *config.Config) error { return commandMap(conf) },
		},
		"mapb": {
			name:"mapb",
			description: "Displays the previous 20 locations queried in map",
			callback: func(conf *config.Config) error { return commandMapB(conf) },
		},
		
	}
}