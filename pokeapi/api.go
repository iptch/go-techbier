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
//    - Use type assertions to return the sprite URL
//    - If anything is missing or of the wrong type, return an error
//
// 2. In `GetAsciiSprite(width int)`:
//    - Use the `image2ascii` package (github.com/zkck/image2ascii)
//    - Generate and return colored ASCII art for display in your Pokédex
//
// 3. Run the tests with `go test ./...` and then the app with `go run .`. Press
//    SPACE on a Pokemon to see your ASCII art!
//
// --------------------------------------------------------------------

package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"image"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
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

	// Note how `v` is of type `any`: an interface with no methods. All structs satisfy this interface.
	var v any
	v = p.Sprites

	for _, key := range keys {
		mapV, ok := v.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("could not go deeper in JSON: not a map")
		}

		v, ok = mapV[key]
		if !ok {
			return "", fmt.Errorf("key not found: %s", key)
		}
	}

	// TODO: Assert `v` is now a string creating a `spriteUrl` variable, and return it! Return an error if not a string.
	return "", fmt.Errorf("not implemented")
}

func (p *Pokemon) GetAsciiSprite(width int) (string, error) {
	spriteUrl, err := p.GetSpriteUrl()
	if err != nil {
		return "", err
	}

	response, err := http.Get(spriteUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return "", err
	}

	// --- Task 3 -------------------------------------------------------------
	// We need to convert the Pokemon sprites into ASCII art. We will use the
	// package github.com/zkck/image2ascii. Use `go get <url>` to import this package,
	// and add the URL at the top of the file to import it here.
	//
	// The package has a `DefaultConverter`, which gives a struct to convert the image `img`
	// to ASCII art. Use the `Convert` method for this. Hint: the `height` in the `Convert` method
	// can be set to 0 to maintain the image ratio, and `uint(i)` can be used to convert an int
	// to a uint.

	// TODO: Convert the image to ASCII art!
	return "", fmt.Errorf("not implemented")

}
