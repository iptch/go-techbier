// TASK DESCRIPTION
// ================
// 🧭 Your mission, dear Pokémon Master, is to catch 'em all — programmatically!
//
// The function `GetAllPokemon` must collect every single Pokémon listed in the PokeAPI.
// But beware! The API is paginated. You’ll have to travel from page to page using the
// `next` field in each response — just like following a trail of Pokéballs.
//
// Here's what you need to do:
// 1. Start your journey at `initialPokemonListUrl`.
// 2. Make an HTTP GET request to fetch the current batch of Pokémon.
// 3. Parse the JSON response into a `PokemonListResponse`.
// 4. Add all Pokémon from the `results` field to your Pokédex (i.e., a slice).
// 5. Update the URL to the `next` field and repeat the process until no Pokémon remain.
//
// At the end, return your hard-earned slice of `PokemonRef` — a complete Pokédex.
//
// Good luck, Trainer. The world is waiting. 🌍🔥⚡🌊

package pokeapi

import (
	"encoding/json"
	"net/http"
)

type PokemonRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonListResponse struct {
	Count   int         `json:"count"`
	NextUrl string       `json:"next"`
	Results []PokemonRef `json:"results"`
}

// Don't by shy! Have a look at the response from the API request
const initialPokemonListUrl = "https://pokeapi.co/api/v2/pokemon"

// GetAllPokemon fetches all PokemonRefs from the PokeAPI.
func GetAllPokemon() ([]PokemonRef, error) {
	currentListUrl := initialPokemonListUrl

	// Slice to collect the results
	pokemonRefs := make([]PokemonRef, 0, 0)

	// Think of this as a while loop :)
	for currentListUrl != "" {

		// TODO: Make an HTTP GET request to the current URL

		// TODO: Parse the response body with JSON into a `PokemonListResponse`
		var _ PokemonListResponse

		// TODO: Update stuff here :)

		break // TODO: Remove this when starting your implementation
	}

	return pokemonRefs, nil
}

// GetPokemonCount fetches the number of pokemon in the PokeAPI.
func GetPokemonCount() (int, error) {
	response, err := http.Get(initialPokemonListUrl)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	var pokemonListResponse PokemonListResponse
    if err := json.NewDecoder(response.Body).Decode(&pokemonListResponse); err != nil {
		return 0, err
	}

	return pokemonListResponse.Count, nil
}
