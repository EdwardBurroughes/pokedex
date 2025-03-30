package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/EdBurroughes/pokedex/cmds"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func BuildRepl(cfg *cmds.URLConfig) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		rawText := scanner.Text()
		cleanedInput := cleanInput(rawText)
		if len(cleanedInput) == 0 {
			continue // skip empty lines
		}

		handler, ok := cmds.BuildCommand()[cleanedInput[0]]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			if err := handler.Callback(cfg, cleanedInput[1:]...); err != nil {
				log.Fatal(err)
			}
		}
	}
}
