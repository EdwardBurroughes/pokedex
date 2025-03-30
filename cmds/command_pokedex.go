package cmds

import "fmt"

func CommandPokedex(cfg *URLConfig, args ...string) error {
	pokedexStr := "Your Pokedex:\n"
	for pokemon := range cfg.CatchDB {
		pokedexStr += fmt.Sprintf(" - %s", pokemon)
	}
	fmt.Println(pokedexStr)
	return nil
}
