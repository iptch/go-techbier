---
title: Welcome to Go!
sub_title: The foundation of cloud computing.
author: Zak Cook & Selim Kälin
theme:
  name: dark
---

Agenda
===
- Go Basics
- Standard Types and Syntax
- Structs
- Functions and Pointers
- Error Handling
- Dealing with JSON

<!-- newline -->
- ***You are up! Task 1***
<!-- newline -->

- Loops and Slices
- Packages, Exports, Constants
- Arrays, Slices, Maps
- Methods
- Interfaces

<!-- newline -->
- ***You are up! Task 2***
<!-- newline -->

- Imports
- Go Management Tools

<!-- newline -->
- ***You are up! Task 3***

<!-- end_slide -->

Get Your Fingers Dirty
---

- You get the chance to get your fingers dirty with your first Go project
- After each theory block, we will give you time to mess around in some Go code
- It makes sense if more experienced Gophers sit with less experienced ones

  - Collaboration is very welcome!

- Instructions to get the code skeleton:
<!-- newline -->

```bash
# Install Go: https://go.dev/doc/install

cd ~/Downloads

sudo rm -rf /usr/local/go && tar -C /usr/local -xzf <go-version.tar.gz>
export PATH=$PATH:/usr/local/go/bin

# Check installation
go version

# Get skeleton code
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

<!-- end_slide -->

Standard Types and Syntax
---

- the full language specification can be found at https://go.dev/ref/spec
- types include:
  - boolean
  - numeric
  - string
  - array, slice, and map
  - struct
  - function and interface
  - pointer
  - channel

<!-- end_slide -->

Declaration and Definition Syntax Basics 1
---

```go +line_numbers
package main        // Everything belongs to a package

func main() {       // Braces are used to delimit scopes, e.g. function declarations
    var x int       // Comments start with two slashes
}
```
<!-- end_slide -->

Declaration and Definition Syntax Basics 2
---
```go +line_numbers +exec
package main

import "fmt"        // Other packages can be imported
 
func main() {
    var x int
    fmt.Println(x)  // Uninitialized variables take default value
}
```
<!-- end_slide -->

Declaration and Definition Syntax Basics 3
---
```go +line_numbers +exec
package main

import "fmt"

func main() {
    var x int  
    var myBoolean bool = true   // Variable declaration follows snakeCase syntax
    var (                       // Declaration blocks are delimited by parantheses
        unsignedInteger uint8
        someFloat       float64
        myFirstString   string
    )
    fmt.Printf("%d, %v, %d, %f, %q", x, myBoolean, unsignedInteger, someFloat, myFirstString)
}
```

<!-- end_slide -->

Declaration and Definition Syntax Basics 4
---
```go +line_numbers +exec
package main

import "fmt"

func main() {
    var x int = 5             // Variable assignment, x has to be declared previously
    y := 7.7                  // Or use the shorthand := to infer the type, e.g. float
    fmt.Printf("%d\n%f\n", x, y)
    fmt.Printf("y is of type %T", y)
}
```

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
        fmt.Printf("A: %s\n", x)
    }
    fmt.Printf("B: %s\n", x)
}

```

<!-- end_slide -->

Structs and Visibility 1
---
```go +line_numbers
package main

// Define a new type
type Consultant struct {
    Name      string
    Age       int
    Project   string
    ahvNumber string // Unexported (private) field due to lowercase
}

func main() {}
```

<!-- end_slide -->

Structs and Visibility 2
---
```go +line_numbers +exec
package main

import "fmt"

type Consultant struct {
    Name      string
    Age       int
    Project   string
    ahvNumber string
}

func main() {
    // Initialize fields with order
    host1 := Consultant{"Zak Cook", 27, "BIT CBCD", "756.0001.0002.03"}
    fmt.Println(host1)
}
```

<!-- end_slide -->

Structs and Visibility 3
---
```go +line_numbers +exec {15-22|20}
package main

import "fmt"

type Consultant struct {
    Name      string
    Age       int
    Project   string
    ahvNumber string
}

func main() {
    host1 := Consultant{"Zak Cook", 27, "BIT CBCD", "756.0001.0002.03"}

    // Initialize fields by name
    host2 := Consultant{
        Name: "Selim Kaelin",
        // Forgot how old Selim is 
        Project:   host1.Project,      // Access fields with dot syntax
        ahvNumber: "756.0001.0002.04", // This would not work from another package
    }
    fmt.Println(host2)
}
```

<!-- end_slide -->

Structs and Visibility 4
---
```go +line_numbers +exec
package main

import "fmt"

type Consultant struct {
    Name      string
    Age       int
    Project   string
    ahvNumber string
}

func main() {
///     // This is hidden code to allow execution but not clutter the slide
///     host1 := Consultant{"Zak Cook", 27, "BIT CBCD", "756.0001.0002.03"}
///     host2 := Consultant{
///        Name: "Selim Kaelin",
///        // Forgot how old Selim is 
///        Project:   host1.Project,      // Access fields with dot syntax
///        ahvNumber: "756.0001.0002.04", // This would not work from another package
///    }
    var host3 Consultant
    host3.Name = "Vincent"
    host3.Project = host2.Project
    host3.ahvNumber = "756.0001.0002.05"  // When would this fail?

    fmt.Println(host3)

}
```

<!-- end_slide -->

Functions and Pointers 1
---

To distinguish between functions and methods in Go, we have to look at the
context in which they are defined:

- Functions: standalone procedure, not associated with any object, i.e. a struct

```go +line_numbers +exec 
package main

import "fmt"

// Definition follows pattern:
// func nameOfFunction(argumentsWithTypes) returnType { ... }
func incrementByValue(x int) int {
    return x + 1
}

func main() {
    a := 5
    fmt.Printf("Before: %d\n", a)
    fmt.Printf("After: %d", incrementByValue(a))
}
```

<!-- end_slide -->

Functions and Pointers 2
---

Pointers are defined using the `*` notation and referenced using `&`.

```go +line_numbers +exec {5-8|13}
package main

import "fmt"

// Void return type
func incrementByReference(x *int) {   // Function argument is a pointer to an integer
    *x = *x + 1                       // Dereferencing the pointer to access its value
}

func main() {
    a := 5
    fmt.Printf("Before: %d\n", a)
    incrementByReference(&a)          // Passing by reference
    fmt.Printf("After: %d", a)
}
```

<!-- end_slide -->

Error Handling With if Statements
---

```go +line_numbers +exec {9-10|13-15|18|all}
package main

import (
    "fmt"
    "os"
)

func createFile(filePath string) (int, error) {
    // os.Create() could return an error
    f, err := os.Create(filePath)

    // We handle the potential error like so
    if err != nil {
        return 0, err
    }

    // This is only executed after the function returns
    defer f.Close()

    return fmt.Fprintln(f, "what up")
}

func main() {
    bytesWritten, err := createFile("/tmp/defer.txt")
    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %s\n", err)
        os.Exit(1)
    }
    fmt.Printf("wrote %d bytes.\n", bytesWritten)
}
```

<!-- end_slide -->

All Together Now: Parsing JSON 1
---

Parsing to and from JSON is essential when using web APIs. 
Luckily, the `encoding/json` package in Go provides very user-friendly functionality to do just that.

Given a JSON schema
```json
{
    "full_name": "Peter Parker",
    "age": 22,
    "project": "Spinning nets in NYC"
}
```

We handle parsing in Go accordingly

```go +line_numbers {9-13}
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Consultant struct {
    FullName string `json:"full_name"`
    Age      int    `json:"age"`
    Project  string `json:"project"`
}
```

<!-- end_slide -->

All Together Now: Parsing JSON 2
---
```go +line_numbers +exec
/// package main
/// 
/// import (
///     "encoding/json"
///     "fmt"
///     "os"
/// )
/// 
/// type Consultant struct {
///     FullName string `json:"full_name"`
///     Age      int    `json:"age"`
///     Project  string `json:"project"`
/// }
/// 
func main() {
    f, err := os.Open("./big_p.json")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // We will parse the JSON into a variable of type Consultant
    var consultant Consultant

    // The not yet initialized variable is passed by reference
    err = json.NewDecoder(f).Decode(&consultant)
    if err != nil {
        panic(err)
    }

    fmt.Println(consultant)
}
```

<!-- end_slide -->

Task 1
===

**Now you are up!**

Open our git repository and check out the branch `tasks/1`.

Look around the project and check out the file `pokeapi/api.go`.

You will find instructions in the code.

We will continue in about _20 minutes_. 
The next slide contains some details about for loops and slices, which you will need to solve task 1b.

<!-- end_slide -->

For Loops and Slices 1
---
Sticking to its philosophy of not overloading the language with too much functionality (looking at you C++, Python), Go supports `for` loops and `for` loops only.  
Usually, a loop is used to iterate through a *slice*, which is Go's equivalent of arrays in other languages.


```go +line_numbers 
package main

func main() {
    var numbers = make([]int, 0)   // Create a slice of type integers and length 0

    // Use a for loop to fill the slice with values
    for i := 0; i < 3; i++ {
        numbers = append(numbers, i)
    }
}
```

<!-- end_slide -->

For Loops and Slices 2
---
```go +line_numbers 
package main

func main() {
    var numbers = make([]int, 0)

///    for i := 0; i < 3; i++ {
///        numbers = append(numbers, i)
///    }
/// 
    // To imitate a while loop
    i := 3
    for i != 5 {
        numbers = append(numbers, i)
        i += 1
    }
}
```
<!-- end_slide -->

For Loops and Slices 3
---
```go +line_numbers 
package main

func main() {
    var numbers = make([]int, 0)

///    for i := 0; i < 3; i++ {
///        numbers = append(numbers, i)
///    }
/// 
///    i := 3
///    for i != 5 {
///        numbers = append(numbers, i)
///        i += 1
///    }
/// 
    // Add many elements at once
    moreNumbers := []int{5, 6, 7}
    numbers = append(numbers, moreNumbers...)
}
```
<!-- end_slide -->

For Loops and Slices 4
---
```go +line_numbers +exec {10-13}
package main

import "fmt"

func main() {
    var numbers = make([]int, 0)

    // ...

///    for i := 0; i < 3; i++ {
///        numbers = append(numbers, i)
///    }
/// 
///    i := 3
///    for i != 5 {
///        numbers = append(numbers, i)
///        i += 1
///    }
/// 
///    moreNumbers := []int{5, 6, 7}
///    numbers = append(numbers, moreNumbers...)
/// 
    // Finally, we print the slice using the range form
    for index, value := range numbers {
        fmt.Printf("Value at index %d: %d\n", index, value)
    }
}
```
<!-- end_slide -->

Packages, Exports, and Constants Syntax Basics
---
Let's compare some features in Go with another popular language, Java.

<!-- column_layout: [1, 1] -->
<!-- column: 0 -->
Everything in go belongs to a package
```go
package main
```

Lowercase letter objects are NOT exported to other packages
```go
var numberInMainPackage = 42
```

Uppercase names are exported
```go
var ExportedString = "Interpackagenal string"
```

Constants are declared like so
```go
const Pi float64 = 3.1415926
```

<!-- column: 1 -->
Similar in Java:
```java
package ch.ipt.ch;
```
<!-- newlines: 2 -->
```java
private static int numberInMainPackage = 42;
```
<!-- newlines: 2 -->
```java
public static String publicString = "Get it?";
```
<!-- newlines: 2 -->
```java
public static final float PI = 3.1415926
```

<!-- end_slide -->

Methods and Their Syntax
---

To distinguish between functions and methods in Go, we have to look at the
context in which they are defined:

Methods:
- like a function but contains a *receiver*, which specifies what type the
  method belongs to
- receiver can be any type, but in most cases it is a struct or pointer to a
    struct

```go
// Unexported method with a return value of type string, c Consultant is the receiver object
func (c Consultant) getProject() string {
    return c.ahvNumber
}
```

<!-- end_slide -->

Interfaces
---

Interfaces specify a list of methods. A type set defined by an interface is the
type set that implements all of those methods.

> **IMPORTANT** In Go, interfaces are implemented **implicitly**! There is no
> explicit declaration of intent, such as the keyword `implements`.

Syntax:

```go +line_numbers
// If it quacks like a duck it is a duck
type Duck interface {
    Quack()
    Swim(distance int)
}
```

We can implement this interface by defining all its methods, i.e. `Quack()`, for a receiver of our choice.
The method definition must be identical to the definition in the interface, i.e. same function arguments.  
Consequently, it is easy to (accidentally) make a duck out of a goose:

```go +line_numbers
type Goose struct {}

func (g Goose) Quack() {
    fmt.Println("Goose Quack!")
}

func (g Goose) Swim(distance int) {
    fmt.Printf("I swam for %d meters!", distance)
}
```

<!-- end_slide -->

Task 2
===

**Now you are up!**

Open our git repository and check out the branch `tasks/2`.

Look around the project and check out the file `ui/item.go`.

You will find instructions in the code.

We will continue in about _20 minutes_.

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

Maps and "comma ok" Notation 1
---

A `map` is Go's hash table equivalent. It is similar to a `dict` in Python.

```go +line_numbers
package main

func main() {
    // Declare according to: map[KeyType]ValueType
    myFirstMap := make(map[string]int)

    myFirstMap["key1"] = 7
    myFirstMap["key2"] = 13
}
```

<!-- end_slide -->

Maps and "comma ok" Notation 2
---

```go +line_numbers
package main

func main() {
    // Instantiate a key-value pair right away
    mySecondMap := map[string]int{
        "key1": 7,
    }
}
```

<!-- end_slide -->

Maps and "comma ok" Notation 3
---

```go +line_numbers +exec
package main

import "fmt"

func main() {
///    myFirstMap := make(map[string]int)
/// 
///    myFirstMap["key1"] = 7
///    myFirstMap["key2"] = 13
/// 
///    mySecondMap := map[string]int{
///        "key1": 7,
///    }
/// 
    for k2, v2 := range mySecondMap {
        // Handle the possibily of k2 not existing in the map using the comma-ok method
        v1, ok := myFirstMap[k2]
        if ok && v1 == v2 {
            fmt.Printf("%s is present and equal in both maps\n", k2)
        }
    }
}
```

<!-- end_slide -->

The Any Type
---
If we don't know yet what a certain variable type will be, or we don't care, we can use two different ways to declare a variable of any type.

```go +line_numbers {5|10}
package main

import "fmt"

func dealsWithAnything(literallyAnything any) {
    fmt.Println("I don't care what you pass to me!")
}

func main() {
    var canBeAnything interface{}
}
```
<!-- end_slide -->

Type Assertions
---
We have seen that Go can infer types when we define variables using the `:=` notation.
Sometimes, we want to ensure that we are dealing with a specific type, and if need be, we want to cast whatever input we get to that type.  
This can be solved using *type assertions* in Go.

```go +line_numbers +exec {10|14|all}
package main

import "fmt"

func main() {
    var canBeAnything interface{}
    canBeAnything = "a string"

    // Type assertion: we are telling Go "this is definitely a string, convert to one"
    ofTypeString := canBeAnything.(string)
    fmt.Println(ofTypeString)

    // Comma ok notation possible
    _, ok := canBeAnything.(int)
    if !ok {
        fmt.Println("wasn't an int")
    }
}
```

<!-- end_slide -->

Go Management Tools
---

- Just like its package management, Go offers very capable management tools
  - `go fmt` for code formatting
  - `go mod`, `go get`, and `go install` for module and dependency management
  - `go test` for testing

<!-- end_slide -->

Task 3
===

**Now you are up!**

Open our git repository and check out the branch `tasks/3`.

Look around the project and check out the file `pokeapi/api.go`.

You will find instructions in the code. We have added a test in `pokeapi/api_test.go`,
check it out and run it with `go test ./...`.

We will continue in about _20 minutes_.

<!-- end_slide -->

Channels and Goroutines 1
---
This is a rather advanced topic that covers concurrency and concurrent programming.
Go is very capable in terms of concurrent programming through its concepts of `channels` and `goroutines`.

- Channels allow data/message passing between concurrent threads of a Go programm 
- Goroutines allow us to run certain functions or code blocks safely in a separate thread
- Goroutines are spawned using the `go` keyword

```go +line_numbers
package main

// Pass string data between threads through the channel messages
func sendMessages(messages chan string) {
    messages <- "ping"
    messages <- "pong"
    close(messages)
}

func main() {
    // ...
}
```

<!-- end_slide -->

Channels and Goroutines 2
---
```go +line_numbers {14-20}
package main

import "fmt"

func sendMessages(messages chan string) {
    messages <- "ping"
    messages <- "pong"
    close(messages)
}

func main() {
    messages := make(chan string)

    // Starts in new thread
    go sendMessages(messages)

    // Wait for messages on the channel until it is closed
    for msg := range messages {
        fmt.Println(msg)
    }
}
```

<!-- end_slide -->

Bonus Tasks
===

Wow! You have come a long ways.

If you are still wanting to play around more, have a look at the branch
`tasks/bonus`.

You will want to start in `pokeapi/api.go`.

<!-- end_slide -->

Useful Resources
===

- Go official documentation: https://go.dev/doc/
- Effective Go (must-read): https://go.dev/doc/effective\_go
- awesome-go: https://github.com/avelino/awesome-go
