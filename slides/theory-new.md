---
title: Go Trainer Bootcamp
sub_title: Build your Pokédex. Master concurrency. Catch 'em all with Go!
author: Zak, Selim, Vincent & Pickachu
theme:
  name: dark
---

Get Your PokeBalls Dirty
===

Welcome to your mission: **help Prof. Oak build the ultimate digital Pokédex!**

- After each theory block, you will work on a practical coding challenge.
- More experienced Gophers are encouraged to team up with newcomers.

  - Collaboration is key to success!

- Setup Instructions:

```bash
# Install Go: https://go.dev/doc/install

cd ~/Downloads

sudo rm -rf /usr/local/go && tar -C /usr/local -xzf <go-version.tar.gz>
export PATH=$PATH:/usr/local/go/bin

# Check installation
go version

# Get the code skeleton
git clone https://github.com/iptch/go-techbier-2024.git
```

<!-- end_slide -->

Go Basics
---

About Go ...

- created in 2009 by R. Griesemer, R. Pike, and K. Thompson at Google
- statically typed and compiled, including to standalone binaries
- features memory safety, garbage collection, structural typing
- built for simplicity and efficiency, i.e. no classes or inheritance
- built-in support for concurrency through `goroutines` and `channels`
- powerful standard library
- great and supportive tooling, e.g. `go test`
- backbone of cloud technology like Kubernetes

> ✨ In our Pokédex project, Go helps us make fast, portable code that runs anywhere!

<!-- end_slide -->

Standard Types and Syntax
---

- The full language specification can be found at https://go.dev/ref/spec
- Types in Go include:
  - `bool`
  - Numeric types: `int`, `float64`, etc.
  - `string`
  - Composite types: `array`, `slice`, `map`, `struct`
  - `function`, `interface`
  - `pointer`
  - `channel`

> 🧠 We’ll use `struct` to define Pokémon, `map` to look them up by name, and `slice` to hold a dynamic Pokédex.

<!-- end_slide -->

Declaration and Definition Syntax Basics 1
---

```go +line_numbers
package main        // Everything belongs to a package

func main() {       // Braces are used to delimit scopes
    var x int       // Declaring a variable of type int
}
```

> ✨ In Go, `var` declares a variable. All variables must be declared before use.

<!-- end_slide -->

Declaration and Definition Syntax Basics 2
---
```go +line_numbers +exec
package main

import "fmt"        // Importing another package

func main() {
    var x int
    fmt.Println(x)  // Default value of int is 0
}
```

> 📃 Variables have default zero-values if not explicitly initialized.

<!-- end_slide -->

Declaration and Definition Syntax Basics 3
---
```go +line_numbers +exec
package main

import "fmt"

func main() {
    var x int
    var myBoolean bool = true
    var (
        unsignedInteger uint8
        someFloat       float64
        myFirstString   string
    )
    fmt.Printf("%d, %v, %d, %f, %q", x, myBoolean, unsignedInteger, someFloat, myFirstString)
}
```


> 💡 Declaration blocks are a clean way to group related variables.

<!-- end_slide -->

Declaration and Definition Syntax Basics 4
---
```go +line_numbers +exec
package main

import "fmt"

func main() {
    var x int = 5
    y := 7.7
    fmt.Printf("%d\n%f\n", x, y)
    fmt.Printf("y is of type %T", y)
}
```

> 📄 Use the shorthand `:=` inside functions, including `main`, to declare and initialize in one line.

<!-- end_slide -->

Pop Quiz
---
```go +line_numbers +exec
package main

import "fmt"

func main() {
    x := "who"
    {
        x := "can guess"
        x = "this variable?"
        fmt.Printf("A: %s\n", x)   // What is the value of x for A?
    }
    fmt.Printf("B: %s\n", x)       // What is the value of x for B?
}
```
> 🔄 Each `{}` block introduces a new scope. Variable `x` inside the block is **not** the same as outside.

<!-- end_slide -->

Structs and Visibility 1
---
```go +line_numbers
package main

// Define a new type for our Pokédex entries
// Structs group related data together
// Fields starting with a lowercase letter are private

type Pokemon struct {
    Name     string
    Type     string
    Level    int
    pokedexID string // unexported field
}

func main() {}
```

> 🧬 Think of a `struct` as a blueprint for a Pokémon entry.

<!-- end_slide -->

Structs and Visibility 2
---
```go +line_numbers +exec
package main

import "fmt"

type Pokemon struct {
    Name     string
    Type     string
    Level    int
    pokedexID string
}

func main() {
    pikachu := Pokemon{"Pikachu", "Electric", 25, "#025"}
    fmt.Println(pikachu)
}
```

> 🐭 We just created our first Pokémon entry! Struct values can be printed directly.

<!-- end_slide -->

Structs and Visibility 3
---
```go +line_numbers +exec {15-22}
package main

import "fmt"

type Pokemon struct {
    Name     string
    Type     string
    Level    int
    pokedexID string
}

func main() {
    pikachu := Pokemon{"Pikachu", "Electric", 25, "#025"}
    fmt.Println(pikachu)
    bulbasaur := Pokemon{
        Name: "Bulbasaur",
        Type: "Grass",
        Level: 12,
        pokedexID: "#001",
    }
    fmt.Println(bulbasaur)
}
```
> 🧪 Named field initialization makes your code more readable and flexible.

<!-- end_slide -->

Structs and Visibility 4
---
```go +line_numbers +exec
package main

import "fmt"

type Pokemon struct {
    Name     string
    Type     string
    Level    int
    pokedexID string
}

func main() {
    var charmander Pokemon
    charmander.Name = "Charmander"
    charmander.Type = "Fire"
    charmander.Level = 18
    charmander.pokedexID = "#004"

    fmt.Println(charmander)
}
```

> 🧯 You can also set struct fields one by one after declaration.

<!-- end_slide -->

Error Handling: Caught an Error!
---

Just like a Pokéball can fail to catch a Pokémon, some operations in Go can fail too — like creating files or opening data.

Go encourages you to **check errors explicitly** using the `if err != nil` pattern.

```go +line_numbers
func throwPokeball(pokemon string) error {
    p, err := pokeball.Catch(pokemon)
    if err != nil {
        return err // Oh no! The Pokéball missed!
    }
    defer p.Close() // Defer the execution of this code to **whenever the function exits**

    fmt.Printf("You caught a Pokemon!")
    return nil // Returning a nil error means success
}
```

> ⚠️ Always check your Pokéballs... err, errors!

<!-- end_slide -->

Parsing Wild Pokémon: JSON 1
---

Oak’s assistant sends Pokémon data in **JSON** format. Luckily, the encoding/json package in Go provides very user-friendly functionality to do just that.

Wild Pokémon report:
```json
{
    "full_name": "Charmander",
    "age": 5,
    "project": "Training to evolve"
}
```

We’ll define a matching struct like this:

```go +line_numbers
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Pokemon struct {
    FullName string `json:"full_name"`
    Age      int    `json:"age"`
    Project  string `json:"project"`
}
```

> 📦 Struct tags like `json:"..."` map JSON keys to Go fields.

<!-- end_slide -->

Parsing Wild Pokémon: JSON 2
---
```go +line_numbers +exec {15-30}
package main
 
import (
     "encoding/json"
     "fmt"
     "os"
)
 
type Pokemon struct {
     FullName string `json:"full_name"`
     Age      int    `json:"age"`
     Project  string `json:"project"`
}
 
func main() {
    f, err := os.Open("./charmander_report.json")
    if err != nil {
        panic(err) // Team Rocket sabotaged the file!
    }
    defer f.Close()

    var mon Pokemon

    err = json.NewDecoder(f).Decode(&mon)
    if err != nil {
        panic(err)
    }

    fmt.Println(mon)
}
```

> 🧾 `json.NewDecoder(f).Decode(&target)` is like scanning a Pokédex entry into our program.

<!-- end_slide -->

For Loops and Slices 1
---

In Go, the only looping keyword is `for` — and it's all you need!
Just like scanning Pokémon one by one, we can loop through a list (slice) of entries.

```go +line_numbers +exec {8-10}
package main

import "fmt"

func main() {
    var pokedex = make([]string, 0)

    for i := 1; i <= 3; i++ {
        pokedex = append(pokedex, fmt.Sprintf("Pokémon #%d", i))
    }
    fmt.Printf("Our current Pokedex: %v\n", pokedex)
}
```

> 🔁 Use `append` to add to a slice — our Pokédex is growing!

<!-- end_slide -->

For Loops and Slices 2
---

You can mimic a `while` loop using `for` in Go — great for catching Pokémon until you run out of Pokéballs!

```go +line_numbers +exec {8-12}
package main

import "fmt"

func main() {
    var pokedex = make([]string, 0)

    pokeballs := 3
    for pokeballs > 0 {
        pokedex = append(pokedex, fmt.Sprintf("Caught #%d", pokeballs))
        pokeballs--
    }
    fmt.Printf("Our current Pokedex: %v\n", pokedex)
}
```

> 🥎 You can use `for` as a `while` — Go keeps things simple.

<!-- end_slide -->

For Loops and Slices 3
---

Want to add many Pokémon at once? Use `append(...slice...)`!

```go +line_numbers +exec {6-8}
package main

import "fmt"

func main() {
    pokedex := []string{"Pikachu", "Charmander"}
    more := []string{"Bulbasaur", "Squirtle"}
    pokedex = append(pokedex, more...)
    fmt.Printf("Our current Pokedex: %v\n", pokedex)
}
```

> 📚 This is like importing a batch of Pokémon entries into your Pokédex.

<!-- end_slide -->

For Loops and Slices 4
---

Loop over a slice using `range` — get both the index and the value!

```go +line_numbers +exec {7-10}
package main

import "fmt"

func main() {
    pokedex := []string{"Pikachu", "Charmander", "Bulbasaur"}

    for index, name := range pokedex {
        fmt.Printf("%d: %s\n", index, name)
    }
}
```

> 📖 Use `range` to loop over your slice like flipping through Pokédex pages.

<!-- end_slide -->

Task 1
===

🔍 Your task is to help Professor Oak fill his Pokédex with Pokémon retrieved from the PokéAPI.

Your mission:

- Implement a function that fetches data from the PokéAPI
- Parse the JSON into Go structs
- Store Pokémon entries in a slice
- Print out the names to verify it works

Open our git repository and check out the branch `tasks/1`.

Look around the project and check out the file `pokeapi/api.go`.

You will find instructions in the code.

We’ll regroup in **20 minutes**.  
The next slide contains some details about for loops and slices, which you will need to solve task 1b.

<!-- end_slide -->

Packages, Exports, and Constants
---

Every Go file starts by declaring its package — just like assigning a Pokémon to a region.

```go
package kanto
```

Let’s explore how exporting works:

```go
// Only exported if the name starts with a capital letter
var PokeballCount = 42

// Can this be used outside the current package?
var rareCandy = 3

const MaxLevel = 100    // Exported or not?
```

> 🧳 Capitalized = Public (exported). Lowercase = Private (unexported).

<!-- end_slide -->

Functions
---

Let’s now look at how Go uses **functions**. A function in Go is a standalone unit of logic that can take input arguments and return values.

```go +line_numbers +exec
package main

import "fmt"

func increaseLevel(level int) int {
    return level + 1
}

func main() {
    original := 5
    newLevel := increaseLevel(original)
    fmt.Printf("Before: %d, After: %d", original, newLevel)
}
```

> 🧪 This passed a **copy** of `level`. To change the original, we need **pointers** (see backup slide for more information).

<!-- end_slide -->
Methods and Receivers
---

To distinguish between functions and methods in Go, we have to look at the context in which they are defined:

Methods:

- like a function but contains a receiver, which specifies what type the method belongs to
- receiver can be any type, but in most cases it is a struct or pointer to a struct

```go
package main

import "fmt"

type Pokemon struct {
    Name string
    Level int
}

// A method with value receiver (copy)
func (p Pokemon) Speak() {
    fmt.Printf("%s says hello!\n", p.Name)
}

// A method with pointer receiver (can modify)
func (p *Pokemon) Train() {
    p.Level++
}
```

> 🎓 Use pointer receivers when your method should **mutate** the struct (like training a Pokémon!).

<!-- end_slide -->

Interfaces: What's Your Type?
---

In the Pokémon world, each species has abilities. In Go, **interfaces** define what a type *can do* — not what it *is*.

Think of an interface like this:
```go +line_numbers
// If it can attack like a Pokemon, it IS a pokemon!
type Pokemon interface {
    Growl()
    Attack(move string)
}
```

> 🧠 Interfaces describe capabilities — not inheritance!
> 🧠 Go follows the approach of composition over inheritance.
> 🟡 No `implements` keyword needed: implementation is **implicit** in Go.

<!-- end_slide -->

Implementing Pokémons
---

If your type has all the required methods, it **automatically** satisfies the interface. Like how any creature that uses "Attack" and "Growl" is a Pokémon in your team.

```go +line_numbers
package main

import "fmt"

type Pikachu struct{}

func (p Pikachu) Growl() {
    fmt.Println("Pika Pika!")
}

func (p Pikachu) Attack(move string) {
    fmt.Printf("Pikachu used %s!\n", move)
}
```

Here, `Pikachu` satisfies the `Pokemon` interface because it implements all the required methods — no need to declare anything.

<!-- end_slide -->

Interfaces in Battle
---

We can use an interface to build a **Pokédex** of battle-ready creatures:

```go +line_numbers +exec {20-28}
package main

import "fmt"

type Pokemon interface {
    Growl()
    Attack(move string)
}

type Pikachu struct{}

func (p Pikachu) Growl() {
    fmt.Println("Pika Pika!")
}

func (p Pikachu) Attack(move string) {
    fmt.Printf("Pikachu used %s!\n", move)
}

func Battle(p Pokemon) {
    p.Growl()
    p.Attack("Thunderbolt")
}

func main() {
    pikachu := Pikachu{}
    Battle(pikachu)
}
```

> 🎮 Pass your Pokémon around as interface values to make your code flexible and extensible.

<!-- end_slide -->


Task 2: Display Pokémon in Style
===

🧾 Now that you’ve fetched some Pokémon, let’s show them off properly!

**Your new task:**
- Head over to `ui/item.go`
- Define a new `Item` type to display Pokémon nicely in a terminal UI
- Implement the required methods (like `Title`, `Description`, etc.) to customize how each entry appears

🎨 Make sure the names, levels, or types are clearly shown

📂 Check the `tasks/2` branch in the repo

We’ll regroup in **20 minutes** to review your stylish Pokédex entries!

> 💡 Hint: Remember your structs and methods!

<!-- end_slide -->

Import Statements
---

- As stated previously, everything in Go belongs to a package, declared by the
  keyword `package`
- Packages are imported using the `import` statement at the beginning of a file
- Imports apply to the entire package, all exported identifiers will become
  available
- Package management is awesome! Look at the following example:

<!-- column_layout: [2, 1] -->
<!-- column: 0 -->
```go
package main

// Let's import multiple packages at once
import (
    "fmt"                                   // Standard library
    "math"                                  // Standard library
    http "net/http"                         // Create an alias called http
    "github.com/charmbracelet/bubbles/list" // External package we will need
)
```
<!-- column: 1 -->
Compare that to Java
```java
import java.util.*;
import java.util.ArrayList;
```

<!-- end_slide -->


Maps
---

A `map` is Go’s built-in way to store key-value pairs. Perfect for Pokémon lookup tables!

```go +line_numbers
package main

func main() {
    pokedex := make(map[string]string)
    pokedex["025"] = "Pikachu"
    pokedex["004"] = "Charmander"
}
```

> 🔍 `map[keyType]valueType` — useful for building fast-access Pokédex indexes.

<!-- end_slide -->

Maps 2
---

You can also directly instantiate a key-value pair!

```go +line_numbers
package main

func main() {
    pokedex := map[string]string{
        "025": "Pikachu",
        "004": "Charmander",
    }
}
```

<!-- end_slide -->

Maps and the Comma OK Idiom
---

Check if a key exists in a map using the `value, ok := map[key]` idiom:

```go +line_numbers +exec {11-18}
package main

import "fmt"

func main() {
    pokedex := map[string]string{
        "025": "Pikachu",
        "004": "Charmander",
    }

    code := "007"
    name, ok := pokedex[code]
    if ok {
        fmt.Printf("%s is in the Pokédex!\n", name)
    } else {
        fmt.Printf("No entry found for code %s.\n", code)
    }
}
```

> ✅ `ok` tells you whether the Pokémon is registered or still hiding in tall grass.

<!-- end_slide -->

Type Assertions: Know Your Pokémon
---

Sometimes, we catch a value of unknown type (`interface{}`), but want to know what it really is — like scanning a Pokémon in the wild!

```go +line_numbers +exec
package main

import "fmt"

func main() {
    var pokeball interface{}
    pokeball = "Eevee"

    // Assert that it's a string
    mon := pokeball.(string)
    fmt.Println(mon)

    // Use comma-ok to avoid runtime panic
    _, ok := pokeball.(int)
    if !ok {
        fmt.Println("That wasn’t a numeric Pokémon!")
    }
}
```

> 🔍 Type assertions let you safely reveal what’s hiding in your Pokéball (i.e. interface).
> 🔍 Remember! Go can infer types when we define variables using the := notation. 

<!-- end_slide -->

Go Tools: Your Trainer Kit
---

Go comes with powerful built-in tools to help you on your dev journey. Here are a few essentials:

- `go mod init` – Start a new Go module (project)
- `go get` – Catch a dependency (like a Poké Ball!)
- `go install` – Install binaries
- `go fmt` – Format your code (style points!)
- `go test` – Run your unit tests

> 🧰 These tools help you build robust, well-organized Go code.

<!-- end_slide -->

Task 3: Gotta Test 'Em All!
===

🧪 You’ve trained your Pokémon... now put them to the test!

**Your final task:**
- Navigate to `pokeapi/api.go`
- Implement and complete the missing function logic
- Explore the test file in `pokeapi/api_test.go`
- Run `go test ./...` to test your Pokédex logic

🎯 Goal: Make sure the Pokédex works as expected and all tests pass.

> 💡 Tests are like Gym battles — prove your code is battle-ready!

📂 Check out the `tasks/3` branch in the repo

We’ll wrap up in **20 minutes** and celebrate with a badge!

<!-- end_slide -->

Channels and Goroutines 1
---

In Go, **channels** and **goroutines** let you handle multiple tasks at once — just like sending out several Pokémon in a double battle!

- **Goroutines** are lightweight threads that run with `go` keyword
- **Channels** let goroutines communicate by passing messages

```go +line_numbers
package main

// Send messages over a channel
func sendMoves(moves chan string) {
    moves <- "Thunderbolt"
    moves <- "Quick Attack"
    close(moves)
}

func main() {
    // ...
}
```

> ⚡ Goroutines make Go ideal for building fast and concurrent apps.

<!-- end_slide -->

Channels and Goroutines 2
---

```go +line_numbers +exec
package main

import "fmt"

func sendMoves(moves chan string) {
    moves <- "Thunderbolt"
    moves <- "Quick Attack"
    close(moves)
}

func main() {
    moves := make(chan string)

    // Starts in a new thread
    go sendMoves(moves)

    // Wait for messages on the channel until closed
    for move := range moves {
        fmt.Println("Pikachu used:", move)
    }
}
```
> 🔁 Use `range` to receive until the channel is closed — like watching Pikachu’s turn-based battle!

<!-- end_slide -->

Bonus Tasks
===

🎉 You've built a Pokédex, parsed wild Pokémon, and learned to test your code.
But maybe you're still thirsty for adventure?

**Try this bonus challenge:**
- Branch: `tasks/bonus`
- File: `pokeapi/api.go`
- Add features like filtering by type, showing top levels, or building a mini CLI Pokédex

> 💡 Stretch tasks are for trainers aiming to be Pokémon Masters 🏆

<!-- end_slide -->

Useful Resources
==

📘 Where to continue your journey:

- Official Go docs: https://go.dev/doc/
- Effective Go: https://go.dev/doc/effective_go
- Curated Go libraries: https://github.com/avelino/awesome-go

> 🧠 Keep learning, keep catching bugs, and keep coding!

<!-- end_slide -->

Thank You from Pikachu!
---
```
quu..__
 $$$b  `---.__
  "$$b        `--.                          ___.---uuudP
   `$$b           `.__.------.__     __.---'      $$$$"              .
     "$b          -'            `-.-'            $$$"              .'|
       ".                                       d$"             _.'  |
         `.   /                              ..."             .'     |
           `./                           ..::-'            _.'       |
            /                         .:::-'            .-'         .'
           :                          ::''\          _.'            |
          .' .-.             .-.           `.      .'               |
          : /'$$|           .@"$\           `.   .'              _.-'
         .'|$u$$|          |$$,$$|           |  <            _.-'
         | `:$$:'          :$$$$$:           `.  `.       .-'
         :                  `"--'             |    `-.     \
        :##.       ==             .###.       `.      `.    `\
        |##:                      :###:        |        >     >
        |#'     `..'`..'          `###'        x:      /     /
         \                                   xXX|     /    ./
          \                                xXXX'|    /   ./
          /`-.                                  `.  /   /
         :    `-  ...........,                   | /  .'
         |         ``:::::::'       .            |<    `.
         |             ```          |           x| \ `.:``.
         |                         .'    /'   xXX|  `:`M`M':.
         |    |                    ;    /:' xXXX'|  -'MMMMM:'
         `.  .'                   :    /:'       |-'MMMM.-'
          |  |                   .'   /'        .'MMM.-'
          `'`'                   :  ,'          |MMM<
            |                     `'            |tbap\
             \                                  :MM.-'
              \                 |              .'\n               \.               `.            /
                /     .:::::::.. :           /
               |     .:::::::::::`.         /
               |   .:::------------\       /
              /   .''               >::'  /
              `',:                 :    .'
                                   `:.:'
```
> ⚡ Thanks for joining our Go Pokédex workshop!

<!-- end_slide -->

Backup
===

<!-- end_slide -->

Pointers 
---

In Go, a pointer is a variable that stores the **memory address** of another variable. When we pass a pointer to a function, we can modify the original value.

```go +line_numbers +exec {5-8|13}
package main

import "fmt"

func boostLevel(level *int) { // *int means pointer to an int
    *level = *level + 1       // *level dereferences the pointer to access the value
}

func main() {
    lvl := 10
    fmt.Printf("Before: %d\n", lvl)
    boostLevel(&lvl)         // Pass the address of lvl
    fmt.Printf("After: %d\n", lvl)
}
```

> 🧠 Use `*` to access and change values via a pointer. Use `&` to get the address of a variable.

> Bonus: If you want a function to modify struct fields directly, you’ll use a **pointer receiver** in a method. We’ll get to that soon!

<!-- end_slide -->
