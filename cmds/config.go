package cmds

import (
	"sync"

	"github.com/EdBurroughes/pokedex/client"
	"github.com/EdBurroughes/pokedex/pokecache"
)

// map[string]Pokemon

type URLConfig struct {
	Next     string
	Previous string
	Client   client.PokeClient
	Cache    *pokecache.Cache
	CatchDB  map[string]client.Pokemon
	mu       sync.RWMutex
}

func (c *URLConfig) SetDirections(nextURL, previousURL string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Next = nextURL
	c.Previous = previousURL
	return nil
}
