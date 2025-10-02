package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ExploreLocation -
func (c *Client) ExploreLocation(area string) (Exploration, error) {
	url := baseURL + "/location-area/" + area

	// Check cache before calling PokeAPI
	dat, exists := c.cache.Get(url)
	if !exists {
		// Call PokeAPI
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Exploration{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return Exploration{}, err
		}
		defer resp.Body.Close()

		cacheDat, err := io.ReadAll(resp.Body)
		if err != nil {
			return Exploration{}, err
		}
		c.cache.Add(url, cacheDat)
		dat = cacheDat
	}

	explorationResp := Exploration{}
	err := json.Unmarshal(dat, &explorationResp)
	if err != nil {
		return Exploration{}, err
	}

	return explorationResp, nil
}
