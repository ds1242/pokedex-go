package pokemap

import (
	"fmt"
	// "io"
	// "log"
	// "net/http"
	"github.com/ds1242/pokedex-go/config"
)



func commandMap(conf *config.Config) error{
	fmt.Println("Get Location Fires")
	GetLocation(conf)
	return nil
}