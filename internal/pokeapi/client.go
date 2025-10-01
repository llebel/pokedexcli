package pokeapi

import (
	"net/http"
	"time"

	"github.com/llebel/pokedexcli/internal/pokecache"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheTtl time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheTtl),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
