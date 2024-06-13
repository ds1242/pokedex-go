package pokemap

import (
	"os"
	"fmt"
	"github.com/ds1242/pokedex-go/config"
)

func commandExit(conf *config.Config) error {
	fmt.Println()
	fmt.Println("Goodbye!")
	fmt.Println()
	os.Exit(0)
	return nil
}