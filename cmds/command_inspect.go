package cmds

import (
	"fmt"
	"strings"

	"github.com/EdBurroughes/pokedex/client"
)

func logPokemonInformation(pokemon client.Pokemon) {
	stats := "Stats:\n"
	for _, stat := range pokemon.Stats {
		stats += fmt.Sprintf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	types := "Types:\n"
	for _, pokeType := range pokemon.Types {
		types += fmt.Sprintf("  -%s\n", pokeType.Type.Name)
	}
	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n%s%s", pokemon.Name, pokemon.Height, pokemon.Weight, stats, types)
}

func CommandInspect(cfg *URLConfig, args ...string) error {
	if len(args) > 1 {
		return fmt.Errorf("more then command has been provided: %s", strings.Join(args, ","))
	}

	name := args[0]
	pokemon, ok := cfg.CatchDB[name]
	if !ok {
		fmt.Printf("Uh oh pokemon %s, hasn't been caught yet\n", name)
		return nil
	}
	logPokemonInformation(pokemon)
	return nil
}
