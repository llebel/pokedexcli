package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemon -
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	// Check cache before calling PokeAPI
	dat, exists := c.cache.Get(url)
	if !exists {
		// Call PokeAPI
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		cacheDat, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}
		c.cache.Add(url, cacheDat)
		dat = cacheDat
	}

	pokemonResp := Pokemon{}
	err := json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}
