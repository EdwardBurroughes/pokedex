package cmds

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/EdBurroughes/pokedex/client"
)

const BaseURL = "https://pokeapi.co/api/v2/location-area"

func logLocationArea(res client.LocationAreaResponse) {
	for _, area := range res.Results {
		fmt.Println(area.Name)
	}
}

func setURLDirection(nextURL, previousURL *string, cfg *URLConfig) error {
	nextURLStr := ""
	if nextURL != nil {
		nextURLStr = *nextURL
	}

	previousURLStr := ""
	if previousURL != nil {
		previousURLStr = *previousURL
	}

	return cfg.SetDirections(nextURLStr, previousURLStr)
}

func handleLocationURL(url string, cfg *URLConfig) error {
	var data client.LocationAreaResponse
	dataBytes, ok := cfg.Cache.Get(url)
	if ok {
		if err := json.NewDecoder(bytes.NewReader(dataBytes)).Decode(&data); err != nil {
			return err
		}
	} else {
		res, err := cfg.Client.GetLocationArea(url)
		if err != nil {
			return err
		}
		data = res
		saveResToCache(cfg.Cache, url, data)
	}

	logLocationArea(data)
	return setURLDirection(data.Next, data.Previous, cfg)
}

func CommandMap(cfg *URLConfig, args ...string) error {
	var url string
	switch {
	case cfg.Next == "" && cfg.Previous == "":
		url = BaseURL
	case cfg.Next != "":
		url = cfg.Next
	case cfg.Next == "" && cfg.Previous != "":
		fmt.Printf("Finished iterating over locations, please use `mapd`")
		return nil
	default:
		return fmt.Errorf("unknown config state")
	}
	if err := handleLocationURL(url, cfg); err != nil {
		return err
	}

	return nil
}

func CommandMapd(cfg *URLConfig, args ...string) error {
	var url string
	switch {
	case cfg.Next == "" && cfg.Previous == "":
		url = BaseURL
	case cfg.Previous != "":
		url = cfg.Previous
	case cfg.Next != "" && cfg.Previous == "":
		fmt.Printf("no previous locations, please use `map`")
		return nil
	default:
		return fmt.Errorf("unknown config state")
	}
	if err := handleLocationURL(url, cfg); err != nil {
		return err
	}

	return nil
}
