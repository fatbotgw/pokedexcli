package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatbotgw/pokedexcli/internal/pokeapi"
	"github.com/fatbotgw/pokedexcli/internal/pokecache"
)

type config struct {
	pCache		*pokecache.Cache
	next		*string
	previous	*string
	pokedex 	map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func runRepl() {
	rand.Seed(time.Now().UnixNano())
	prompt := "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)

	mapConfig := config {
		pCache:		pokecache.NewCache(15 * time.Second),
		next:		nil,
		previous:	nil,
		pokedex:	make(map[string]pokeapi.Pokemon),
	}

	for {
		// display the program prompt
		fmt.Print(prompt)

		// scan user input
		scanner.Scan()

		// run the input through the cleaner
	    cleanedInput := cleanInput(scanner.Text())

        if command, exists := getCommands()[cleanedInput[0]]; exists {
        	arg := cleanedInput[1:]
			err := command.callback(&mapConfig, arg)
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
		"explore": {
			name:		 "explore",
			description: "List all Pokemon in an area",
			callback:	 commandExplore,
		},
		"catch": {
			name:		 "catch",
			description: "Catch the Pokemon",
			callback: 	 commandCatch,
		},
		"inspect": {
			name: 		 "inspect",
			description: "Inpect the Pokemon",
			callback: 	 commandInspect,
		},
		"pokedex": {
			name:		 "pokedex",
			description: "List all Pokemon in the Pokedex",
			callback: 	 commandPokedex,
		},
	}
}
