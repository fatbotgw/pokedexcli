package pokeapi

type Location struct {
    Location struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"location"`

    Name  string `json:"name"` // top-level name of the location-area

    PokemonEncounters []struct {
        Pokemon struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"pokemon"`
    } `json:"pokemon_encounters"`
}
