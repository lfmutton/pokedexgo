package api

import(
	"net/http"
	"time"
	"github.com/lfmutton/pokedex/Internal/Cache"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct{
	httpClient http.Client
	cache pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client{
	return Client{
		cache:  pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

