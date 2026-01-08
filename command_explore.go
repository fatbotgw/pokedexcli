package main

import (
	"fmt"
	"encoding/json"
	"io"
	"net/http"

	"github.com/fatbotgw/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args []string) error {
	// List all Pokemon in a location
	// https://pokeapi.co/api/v2/location-area/{id or name}/
	location := ""
	if len(args) > 0 {
		location = args[0]
	}

	address := "https://pokeapi.co/api/v2/location-area/" + location


	body, exists := cfg.pCache.Get(address)
	if !exists {


		res, err := http.Get(address)
		if err != nil {
			return err
		}

		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			return err
		}

		cfg.pCache.Add(address, body)
	}

	var encounters pokeapi.Location
	if err := json.Unmarshal(body, &encounters); err != nil {
		return err
	}

	for _, encounter := range encounters.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

