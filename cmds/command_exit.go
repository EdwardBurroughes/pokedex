package cmds

import (
	"fmt"
	"os"
)

func CommandExit(cfg *URLConfig, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
