package main
import (
	"github.com/ds1242/pokedex-go/pokemap"
	"github.com/ds1242/pokedex-go/config"
	"github.com/ds1242/pokedex-go/pokecache"
)


func main() {	
	conf := &config.Config{
		// NextURL: "https://pokeapi.co/api/v2/location",
		NextURL: "https://pokeapi.co/api/v2/location?offset=20&limit=20",
	}
	pokemap.StartRepl(conf)
}


