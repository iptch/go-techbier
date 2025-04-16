package pokeapi

import (
	"testing"
)

func TestSpriteUrl(t *testing.T) {
	pokemon := Pokemon{
		Sprites: map[string]interface{}{
			"other": map[string]interface{}{
				"official-artwork": map[string]interface{}{
					"front_default": "https://pokemon.com/dittosprite.png",
				},
			},
		},
	}
	spriteUrl, err := pokemon.GetSpriteUrl()
	if err != nil {
		t.Fatal(err)
	}
	if spriteUrl != "https://pokemon.com/dittosprite.png" {
		t.Fatal("url mismatch")
	}
}
