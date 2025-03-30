package cmds

type cliCommand struct {
	Name        string
	Description string
	Callback    func(cfg *URLConfig, args ...string) error
}

func BuildCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			Name:        "Exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "Help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "Map",
			Description: "Retrieve the Pokemon locations in batch of 20",
			Callback:    CommandMap,
		},
		"mapd": {
			Name:        "Mapd",
			Description: "Retrieve the previous Pokemen locations in a previous batch of 20",
			Callback:    CommandMapd,
		},
		"explore": {
			Name:        "Explore",
			Description: "Explore pokemon at a given location",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "Catch",
			Description: "Try to catch a pokemon, will depend on the pokemons base experience",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "Inspect",
			Description: "Inspect the stats of a pokemon",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "Pokedex",
			Description: "Get all caught pokemon in the pokedex",
			Callback:    CommandPokedex,
		},
	}
}
