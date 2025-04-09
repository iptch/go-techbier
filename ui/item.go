// TASK DESCRIPTION
// ================
// 🧠 Welcome to Task 2, Pokémon Trainer!
//
// You're building a Pokédex UI using the Charm Bubbles list component,
// and you need to make sure each Pokémon in your list knows how to
// introduce itself properly.
//
// Right now, you have a `PokemonItem` type that wraps basic Pokémon info
// (like name and URL). Your goal is to make this type implement the
// `list.DefaultItem` interface from the Bubbles library.
//
// Here's what you need to do:
//
// 1. Implement the following methods on `*PokemonItem`:
//    - `Title() string`         → This should return the Pokémon's name, nicely capitalized
//    - `Description() string`   → You can return an empty string here
//    - `FilterValue() string`   → Return a lowercase version of the name for fuzzy searching
//
// 2. When your code is correct, the line:
//
//       var _ list.DefaultItem = (*PokemonItem)(nil)
//
//    will confirm your implementation is correct (no red squiggles! ✅).
//
// 3. Run your code with `go run`, and search for your favorite Pokemon :)

// 💡 Hints:
// - You can use the `ref pokeapi.PokemonRef` inside your `PokemonItem` struct for easy field access.
// - If you're curious, check out `golang.org/x/text/cases` to make the title prettier.
// - Keep it simple and readable. You're building for clarity, not cleverness!
//
// Ready? Time to code! 🧑‍💻🐱‍🏍🔥

package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/iptch/go-techbier-2024/pokeapi"
)

// PokemonItem wraps a basic Pokémon reference (name + URL) for use in the list UI.
// Your goal: make this type implement the list.DefaultItem interface.
type PokemonItem struct {
	ref pokeapi.PokemonRef
}

// ✅ This line ensures the compiler checks that PokemonItem implements list.DefaultItem.
// It will fail to compile until you've implemented all required methods.
var _ list.DefaultItem = (*PokemonItem)(nil)

// 🧠 Tips:
// - Declare methods starting with `func (p PokemonItem) MyMethod() string { ... }`
// - Capitalize the Pokémon name in Title() for better readability
//   ➜ Check out: golang.org/x/text/cases

// TODO: Add method implementations
