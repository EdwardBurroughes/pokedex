package cmds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/EdBurroughes/pokedex/client"
)

func CommandExplore(cfg *URLConfig, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("more then command has been provided: %s", strings.Join(args, ","))
	}
	area := args[0]

	var data client.ExploreLocationAreaRes
	dataBytes, ok := cfg.Cache.Get(area)
	if ok {
		if err := json.NewDecoder(bytes.NewReader(dataBytes)).Decode(&data); err != nil {
			return err
		}
	} else {
		res, err := cfg.Client.GetExploreArea(area)
		if err != nil {
			return err
		}
		data = res
		saveResToCache(cfg.Cache, area, data)
	}

	for _, pokeObj := range data.PokemonEncounters {
		fmt.Println(pokeObj.Pokemon.Name)
	}
	return nil
}
