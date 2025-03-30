package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokeClient struct {
	HTTPClient *http.Client
}

const BASE_URL = "https://pokeapi.co/api/v2/"

type PokeResponse interface {
	LocationAreaResponse | ExploreLocationAreaRes | Pokemon
}

func getJSON[T PokeResponse](client *PokeClient, url string) (T, error) {
	res, err := client.HTTPClient.Get(url)
	if err != nil {
		var zero T
		return zero, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var zero T
		return zero, fmt.Errorf("non 200 error code: %s", res.Status)
	}
	var data T
	decoder := json.NewDecoder(res.Body)
	decoder.Decode(&data)
	return data, nil
}

func (p *PokeClient) GetLocationArea(url string) (LocationAreaResponse, error) {
	return getJSON[LocationAreaResponse](p, url)
}

func (p *PokeClient) GetExploreArea(area string) (ExploreLocationAreaRes, error) {
	url := fmt.Sprintf("%s/location-area/%s", BASE_URL, area)
	return getJSON[ExploreLocationAreaRes](p, url)
}

func (p *PokeClient) GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", BASE_URL, name)
	return getJSON[Pokemon](p, url)
}

func NewClient() PokeClient {
	return PokeClient{
		HTTPClient: &http.Client{},
	}
}
