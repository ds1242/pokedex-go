package pokemap

import (
	// "fmt"
	"github.com/ds1242/pokedex-go/config"
)



func commandMap(conf *config.Config) error {
	GetLocation(conf)
	return nil
}