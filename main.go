package main

import (
	"time"

	"github.com/EdBurroughes/pokedex/client"
	"github.com/EdBurroughes/pokedex/cmds"
	"github.com/EdBurroughes/pokedex/pokecache"
)

const SECONDS = 10

func main() {
	pokeClient := client.NewClient()
	cache := pokecache.NewCache(
		10 * time.Second,
	)
	cfg := &cmds.URLConfig{
		Client:  pokeClient,
		Cache:   cache,
		CatchDB: make(map[string]client.Pokemon),
	}
	BuildRepl(cfg)
}
