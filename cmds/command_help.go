package cmds

import "fmt"

func CommandHelp(cfg *URLConfig, args ...string) error {
	msg := "Welcome to the Pokedex!\nUsage:"
	helpStr := ""
	for command, handler := range BuildCommand() {
		helpStr += fmt.Sprintf("%s: %s\n", command, handler.Description)
	}
	fmt.Printf("%s\n\n%s", msg, helpStr)

	return nil
}
