package pokemap

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/ds1242/pokedex-go/config"
)

type LocationData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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

	locationData := LocationData{}

	decodErr := json.Unmarshal(body, &locationData)
	if decodErr != nil {
		fmt.Println(decodErr)
	}
	// fmt.Println(locationData)
	fmt.Println(locationData.Next)
	conf.NextURL = locationData.Next
	conf.PreviousURL = locationData.Next
	fmt.Println(conf)

	for _, location := range locationData.Results {
		fmt.Println(location.Name)
	}
	// fmt.Printf("%s", locationData.Results)
	
	// fmt.Printf("%s", body)
	
	return nil
}