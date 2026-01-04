package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func runRepl() {
	prompt := "Pokedex > "
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// display the program prompt
		fmt.Print(prompt)

		// scan user input
		scanner.Scan()

		// run the input through the cleaner
	    cleanedInput := cleanInput(scanner.Text())

        if command, exists := getCommands()[cleanedInput[0]]; exists {
			err := command.callback()
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}