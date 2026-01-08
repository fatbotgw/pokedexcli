package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatbotgw/pokedexcli/internal/pokeapi"
)

// EXAMPLE:
// {"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"}

type LocationResponse struct {
	Count 		int			`json:"count"`
	Next		*string		`json:"next"`
	Previous	*string		`json:"previous"`
	Results		[]pokeapi.Location	`json:"results"`
}

// Displays the names of 20 location areas in the Pokemon world.
//
// Use:
// GET https://pokeapi.co/api/v2/location-area/{id or name}/
//
// Calling without an id or name will return the first 20 locations, to get
// the next (or previous) will require using an offset and a limit. Look at the
// API documentation for more information.

func commandMap(cfg *config, args []string) error {
	// Changed initial address from:
	// address := "https://pokeapi.co/api/v2/location-area/"
	// to:
	address := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	// in order for the caching of the first page to work correctly with
	// the .previous that is used in commandMapb() below.

	if cfg.next != nil {
		address = *cfg.next
	}
	body, exists := cfg.pCache.Get(address)

	if !exists {
		res, err := http.Get(address)
		if err != nil {
			return err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		// if res.StatusCode > 299 {
		// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		// }
		if err != nil {
			return err
		}
		cfg.pCache.Add(address, body)
	}

	var locations LocationResponse
	if err := json.Unmarshal(body, &locations); err != nil {
		return err
	}
	cfg.next = locations.Next
	cfg.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}


func commandMapb(cfg *config, args []string) error {
	if cfg.previous == nil {
		fmt.Println("You're on the first page.")
		return nil
	}
	address := *cfg.previous

	body, exists := cfg.pCache.Get(address)

	if !exists {
		res, err := http.Get(address)
		if err != nil {
			return err
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		// if res.StatusCode > 299 {
		// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		// }
		if err != nil {
			return err
		}
		cfg.pCache.Add(address, body)
	}

	var locations LocationResponse
	if err := json.Unmarshal(body, &locations); err != nil {
		return err
	}
	cfg.next = locations.Next
	cfg.previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Printf("%s\n", location.Name)
	}

	return nil
}
