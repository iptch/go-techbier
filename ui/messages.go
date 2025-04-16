package ui

import "github.com/iptch/go-techbier/pokeapi"

type DownloadCompleted struct {
	// empty
}

type NewPokemon struct {
	Pokemon pokeapi.PokeapiRef[pokeapi.Pokemon]
}
