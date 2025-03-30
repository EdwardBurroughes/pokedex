package cmds

import (
	"bytes"
	"encoding/json"

	"github.com/EdBurroughes/pokedex/client"
	"github.com/EdBurroughes/pokedex/pokecache"
)

func saveResToCache[T client.PokeResponse](cache *pokecache.Cache, key string, data T) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		return err
	}
	cache.Add(key, buf.Bytes())
	return nil

}
