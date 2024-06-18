package main
import (
	"github.com/ds1242/pokedex-go/pokemap"
	"github.com/ds1242/pokedex-go/config"
)


func main() {	
	conf := &config.Config{
		NextURL: "https://pokeapi.co/api/v2/location",
	}
	pokemap.StartRepl(conf)
}


