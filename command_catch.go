package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/fatbotgw/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args []string) error {
	pokeName := ""
	if len(args) > 0 {
		pokeName = args[0]
	}
	fmt.Println("Throwing a Pokeball at " + pokeName + "...")

	// https://pokeapi.co/api/v2/pokemon/{id or name}/
	address := "https://pokeapi.co/api/v2/pokemon/" + pokeName

	res, err := http.Get(address)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	var pokemon pokeapi.Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	maxXP := 500
	pokeXP := pokemon.BaseExperience
	if pokeXP > maxXP {
		maxXP = pokeXP
	}
	chance := 100 - (pokeXP * 100 / maxXP)
	if chance < 5 {
		chance = 5
	}

	roll := rand.Intn(100) // 0..99
	if roll < chance {
		// print "<name> was caught!"
		fmt.Println(pokeName + " was caught!")
		// add to pokedex map
		if cfg.pokedex == nil {
			cfg.pokedex = make(map[string]pokeapi.Pokemon)
		}
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		// print "<name> escaped!"
		fmt.Println(pokeName + " escaped!")
	}

	return nil
}
