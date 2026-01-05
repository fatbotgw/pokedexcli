package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// {"name":"canalave-city-area","url":"https://pokeapi.co/api/v2/location-area/1/"}
type Location struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}

type LocationResponse struct {
	Count 		int			`json:"count"`
	Next		*string		`json:"next"`
	Previous	*string		`json:"previous"`
	Results		[]Location	`json:"results"`
}

// Displays the names of 20 location areas in the Pokemon world.
//
// Use:
// GET https://pokeapi.co/api/v2/location-area/{id or name}/
//
// Calling without an id or name will return the first 20 locations, to get
// the next (or previous) will require using an offset and a limit. Look at the
// API documentation for more information.

func commandMap(cfg *config) error {
	address := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		address = *cfg.next
	}
	res, err := http.Get(address)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	if err != nil {
		return err
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

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("You're on the first page.")
		return nil
	}
	address := *cfg.previous

	res, err := http.Get(address)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	// if res.StatusCode > 299 {
	// 	log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	// }
	if err != nil {
		return err
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
