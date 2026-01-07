package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatbotgw/pokedexcli/internal/pokecache"
)

type config struct {
	pCache		*pokecache.Cache
	next		*string
	previous	*string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func runRepl() {
	prompt := "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)

	mapConfig := config {
		pCache:		pokecache.NewCache(15 * time.Second),
		next:		nil,
		previous:	nil,
	}

	for {
		// display the program prompt
		fmt.Print(prompt)

		// scan user input
		scanner.Scan()

		// run the input through the cleaner
	    cleanedInput := cleanInput(scanner.Text())

        if command, exists := getCommands()[cleanedInput[0]]; exists {
			err := command.callback(&mapConfig)
				if err != nil {
					fmt.Println(err)
				}
			continue
        } else {
            fmt.Println("Unknown command:", cleanedInput[0])
            continue
        }
	}
}

// The purpose of this function will be to split the user's input into "words"
// based on whitespace. It should also lowercase the input and trim any leading
// or trailing whitespace.
func cleanInput(text string) []string {
	retString := strings.Fields(strings.ToLower(text))
    return retString
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:		 "map",
			description: "Displays the names of next 20 location areas in the Pokemon world",
			callback:	 commandMap,
		},
		"mapb": {
			name:		 "mapb",
			description: "Displays the names of previous 20 location areas in the Pokemon world",
			callback: 	 commandMapb,
		},
	}
}