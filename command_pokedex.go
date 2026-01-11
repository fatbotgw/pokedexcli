package main

import (
	"fmt"
)


// EXAMPLE OUTPUT:
// Pokedex > pokedex
// Your Pokedex:
//  - pidgey
//  - caterpie

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.pokedex {
		fmt.Printf(" - %v\n", pokemon)
	}

	return nil
}
