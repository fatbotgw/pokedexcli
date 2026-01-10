package main

import (
	"fmt"
)

// takes the name of a Pokemon and prints the name, height, weight, stats and
// type(s) of the Pokemon
//
// EXAMPLE OUTPUT:
// Pokedex > inspect pidgey
// Name: pidgey
// Height: 3
// Weight: 18
// Stats:
//   -hp: 40
//   -attack: 45
//   -defense: 40
//   -special-attack: 35
//   -special-defense: 35
//   -speed: 56
// Types:
//   - normal
//   - flying

func commandInspect(cfg *config, args []string) error {
	// print Pokemon information
	pokemon, exists := cfg.pokedex[args[0]]
	if exists {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, stat := range pokemon.Types {
			fmt.Printf("  -%s\n", stat.Type.Name)
		}
	} else {
		fmt.Println("You have not caught that Pokemon yet.")
	}

	return nil
}
