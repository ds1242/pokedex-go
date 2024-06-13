package pokemap

import (
	"fmt"

	"github.com/ds1242/pokedex-go/config"
)


func commandHelp(conf *config.Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Commands Available:")
	fmt.Println()
	for _, cmd := range GetCommands(conf) {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}