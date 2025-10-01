package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/llebel/pokedexcli/internal/pokecache"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string, pokeCache *pokecache.Cache) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cache before calling PokeAPI
	dat, exists := pokeCache.Get(url)
	if !exists {
		// Call PokeAPI
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Locations{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Locations{}, err
		}
		defer resp.Body.Close()

		cacheDat, err := io.ReadAll(resp.Body)
		if err != nil {
			return Locations{}, err
		}
		pokeCache.Add(url, cacheDat)
		dat = cacheDat
	}

	locationsResp := Locations{}
	err := json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Locations{}, err
	}

	return locationsResp, nil
}
