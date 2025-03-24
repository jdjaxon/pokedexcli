package api

import (
	"net/http"
	"time"

	"github.com/jdjaxon/pokedexcli/internal/cache"
)

const cacheInterval = 5 * time.Second

type Client struct {
	httpClient http.Client
	reqCache   *cache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		reqCache: cache.NewCache(cacheInterval),
	}
}
