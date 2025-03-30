package cmds

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/EdBurroughes/pokedex/client"
)

const (
	HIGH   = "high"
	MEDIUM = "medium"
	LOW    = "low"
)

func getExperienceCategory(experience int) (string, error) {
	switch {
	case experience > 120:
		return HIGH, nil
	case 100 <= experience && experience < 120:
		return MEDIUM, nil
	case experience < 100:
		return LOW, nil
	default:
		return "", fmt.Errorf("negative experience, %d", experience)
	}
}

func biasedRandom(category string) bool {
	weights := map[string]float64{
		HIGH:   0.1,
		MEDIUM: 0.5,
		LOW:    0.8,
	}

	prob := weights[category]
	return rand.Float64() <= prob
}

func CommandCatch(cfg *URLConfig, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("more then command has been provided: %s", strings.Join(args, ","))
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	var pokemon client.Pokemon
	pokemonDB, ok := cfg.CatchDB[name]
	if ok {
		pokemon = pokemonDB
	} else {
		pokemonRes, err := cfg.Client.GetPokemon(name)
		if err != nil {
			return err
		}
		pokemon = pokemonRes
	}
	cat, err := getExperienceCategory(pokemon.BaseExperience)
	if err != nil {
		return err
	}

	if biasedRandom(cat) {
		fmt.Printf("%s was caught!\n", name)
		cfg.CatchDB[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
