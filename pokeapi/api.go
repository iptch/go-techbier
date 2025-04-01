// PROJECT DESCRIPTION
// ===================
// Welcome, brave developer-trainer! 🧑‍💻⚡🐉
//
// You're working on the backend of a Pokédex app powered by the PokeAPI.
// Your goal is to fetch Pokémon data, explore their abilities, and even generate
// beautiful ASCII sprite art to display in the UI.
//
// The `pokeapi` package provides a simple and extendable client for accessing the PokeAPI,
// fetching Pokémon references, detailed stats, types, and sprite images.
//
// So far, you’ve got types like:
//
// - `PokemonRef`   → A basic reference to a Pokémon (name + URL)
// - `Pokemon`      → Full details about a Pokémon (types, stats, and sprites)
// - `PokemonTypeRef` and `PokemonStatRef` → Subobjects within a Pokémon
//
// You’ve also got functions that:
//
// ✅ Fetch a list of Pokémon using pagination (`GetAllPokemon(n int)`)
// ✅ Fetch full details of a Pokémon via `.Get()` on a `PokemonRef`
//
// --------------------------------------------------------------------
// 🧠 TASK 3: Working with Sprites & ASCII Art
// --------------------------------------------------------------------
//
// Your mission now is to extract the Pokémon's official artwork URL
// from the `Sprites` map and convert that image into ASCII art!
//
// Specifically:
//
// 1. In `GetSpriteUrl()`:
//    - Traverse the `Sprites` map to find the URL under:
//      sprites → other → official-artwork → front_default
//    - Use type assertions to navigate safely through each layer
//    - If anything is missing or of the wrong type, return an error
//
// 2. In `GetAsciiSprite(width int)`:
//    - Use the `ascii-image-converter` package (github.com/TheZoraiz/ascii-image-converter/aic_package)
//    - Fetch the sprite URL using your `GetSpriteUrl()` function
//    - Use the `DefaultFlags()` and `Convert()` methods from the package
//    - Generate and return colored ASCII art for display in your Pokédex
//
// --------------------------------------------------------------------

package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonTypeRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonStatRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonRef struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokemonList struct {
	Results []PokemonRef `json:"results"`
	NextUrl string       `json:"next"`
}

type PokemonStat struct {
	// no fields used
}

type PokemonType struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Name  string `json:"name"`
	Types []struct {
		Slot int            `json:"slot"`
		Type PokemonTypeRef `json:"type"`
	} `json:"types"`
	Stats []struct {
		BaseStat int            `json:"base_stat"`
		Stat     PokemonStatRef `json:"stat"`
	} `json:"stats"`
	Sprites map[string]interface{} `json:"sprites"`
}

// GetAllPokemon reads all available Pokémon from the pokeapi incrementally.
// A GET on the url provided returns a list of results and a next URL to perform
// another GET request on for another set of Pokémon.
func GetAllPokemon(n int) ([]PokemonRef, error) {
	results := make([]PokemonRef, 0)

	url := "https://pokeapi.co/api/v2/pokemon"

	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		var pokemonList PokemonList
		err = json.NewDecoder(resp.Body).Decode(&pokemonList)
		if err != nil {
			return nil, err
		}

		results = append(results, pokemonList.Results...)
		if len(results) >= n {
			return results[:n], nil
		}

		url = pokemonList.NextUrl
	}

	return results, nil
}

func (p PokemonRef) Get() (*Pokemon, error) {
	resp, err := http.Get(p.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return nil, err
	}

	return &pokemon, nil
}

func (p *Pokemon) GetSpriteUrl() (string, error) {
	keys := []string{"other", "official-artwork", "front_default"}

	spritesMap := p.Sprites

	var spritesUrl string
	// ### Task 3 ###
	// You will need to use the index returned by `range`, so replace the placeholder
	for _, key := range keys {

		// ### Task 3 ###
		// You will need to use the value from spritesMap, so replace the placeholder
		_, ok := spritesMap[key]
		if !ok {
			return "", fmt.Errorf("key not found: %s", key)
		}

	// --- Task 3 -------------------------------------------------------------
	// Your mission is to extract the front sprite URL from the nested sprite data.
	//
	// The desired path in the JSON response looks like this:
	//     sprites → other → official-artwork → front_default
	//
	// To achieve this:
	//   1. Iterate over the `keys` slice above
	//   2. For each key:
	//      a. Check if the current `spritesMap` contains that key
	//      b. If not, return an error (the data is incomplete)
	//      c. If it exists:
	//         - If we're at an intermediate key, assert it's a `map[string]interface{}`
	//           and continue traversal deeper (update spritesMap)
	//         - If it's the last key (`front_default`), assert it's a `string` and assign
	//           it to `spritesUrl`
	//         - If anything fails along the way (type mismatch, missing key), return an error
	//
	// Pro Tip: Look at actual JSON responses on https://pokeapi.co to get a feel for the structure
	// -------------------------------------------------------------------------

	}

	return spritesUrl, nil
}

func (p *Pokemon) GetAsciiSprite(width int) (string, error) {
	spriteUrl, err := p.GetSpriteUrl()
	if err != nil {
		return "", err
	}

	// --- Task 3 -------------------------------------------------------------
	// We need to convert the Pokemon sprites into ASCII art. We will use the
	// package github.com/TheZoraiz/ascii-image-converter/aic_package.
	//
	// Add the necessary import statements at the top of the file and use the
	// imported package to create an ASCII sprite for our Pokédex.

	// Uncomment these lines once you are ready
	// flags := aic_package.DefaultFlags()
	// flags.Width = width
	// flags.Colored = true

	// Finally, wou will need a second function from the imported package, called Convert().
	// The Convert function takes two parameters: a sprite URL and corresponding flags.
	// Make sure you adjust the return statement correctly.

	return spriteUrl, fmt.Errorf("not implemented yet")

}
