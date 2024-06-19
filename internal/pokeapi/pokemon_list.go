package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonInfo := Pokemon{}
		err := json.Unmarshal(val, &pokemonInfo)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonInfo := Pokemon{}
	err = json.Unmarshal(dat, &pokemonInfo)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemonInfo, nil
}