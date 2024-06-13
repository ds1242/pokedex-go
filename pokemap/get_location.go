package pokemap

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/ds1242/pokedex-go/config"
)

func GetLocation(conf *config.Config) error {
	res, err := http.Get(conf.NextURL)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
	
	return nil
}