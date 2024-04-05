# Learning Go Through `Cards` Project

---

- [Variables](#variables)
  - [Basic Variable Declaration Format](#basic-variable-declaration-format)
  - [Walrus Variable Declaration Format](#walrus-variable-declaration-format)
  - [Re-Assigning Values To Variables](#re-assigning-values-to-variables)
- [Functions](#functions)
  - [Basic Function Declaration Format](#basic-function-declaration-format)
  - [Tuple-Like Assignement And Usage](#tuple-like-assignement-and-usage)
- [Arrays \& Slices](#arrays--slices)
  - [Array](#array)
    - [Accessing Array Element](#accessing-array-element)
  - [Slice](#slice)
    - [Accessing Slice Element](#accessing-slice-element)
    - [Selecting A Subset/Range Of Slice](#selecting-a-subsetrange-of-slice)
    - [Appending New Elements To A Slice](#appending-new-elements-to-a-slice)
  - [Iteration With `for`-Loops](#iteration-with-for-loops)
    - [Iteration Over Finite Sets](#iteration-over-finite-sets)
    - [Iteration Over Infinite Sets](#iteration-over-infinite-sets)
    - [Iteration in C-style](#iteration-in-c-style)
    - [`break`, `continue`, And `do-while`](#break-continue-and-do-while)
- [Types \& Receiver functions](#types--receiver-functions)
  - [Types](#types)
  - [Initializer Function](#initializer-function)
  - [Receiver Function](#receiver-function)
    - [General Declaration Format Of A Receiver Function](#general-declaration-format-of-a-receiver-function)
    - [Example Of A Receiver Function](#example-of-a-receiver-function)
  - [Using Types](#using-types)
- [Writing To File](#writing-to-file)
- [Reading From File](#reading-from-file)
  - [Error Handling](#error-handling)
- [Unit Test](#unit-test)
  - [How Do We Know What To Test?](#how-do-we-know-what-to-test)
  - [Testing In Go](#testing-in-go)

---

In this review, we are implementing [a system of cards and deck](./system-design.md), and review the following concepts:

- Variables
- Functions
- Arrays & Slices
- Iteration with `for`-loops
- Types & Receiver Functions
- Writing To File
- Reading From File
- Error Handling
- Unit Test

## Variables

- Variables can be initialized inside or outside of a function
- **But can only be assigned a value inside a function**
- Go uses the `var` keyword to declare variables
- *Variables are typed*: Go is a statically-typed language
- **Every declared variables must be used**

### Basic Variable Declaration Format

```go
// Declare
// var <name> <type>
var myCard string

// Initialize
// <name> = <value>
myCard = "Ace of Spade"

// Declare and Initialize
// var <name> <type> = <value>
var yourCard string = "Jack of Heart"
```

- `var`
  - Keyword to create a new variable
  - **Every declared variables must be used**
- `myCard`
  - Name of the variable
- `string`
  - Data type of the variable
  - *Go is a statically-typed language*
  - The variable type follows the variable name
- Go fundamental types
  - `string`
  - `bool`
  - `int`
  - `float64`
- We can also declare only, then assigned a value later
  - When using this approach, Go assigns the *null*-equivalent default value of the type to the declared variable

```go
// Declare variable:
// Default value for string => ""
var someCard string

// Declare variable:
// Default value for int => 0
var someInt int

// Assign value to variables later
someCard = "5 of Heart"
someInt = 1001
```

### Walrus Variable Declaration Format

- Go can also automatically *infer* the variable type from the assigned value
  - We use `:=` and omit the `var` keyword
  - **Only use `:=` when declaring a new variable *WITH* initialization *AND* type inference**

```go
// These are equivalent to the above declarations
someCard := "5 of Heart"   // type: string
someInt := 1001            // type: int
```

### Re-Assigning Values To Variables

- Obviously, we can re-assign value to any variable
- We can only assign value of the same type
- **However, make sure to use `=` instead of `:=` when re-assigning**
  - `:=` is only used for initializing with type inference
  - `=` is used for all successive re-assignments
- **Make sure that the type of the value matches the declared type of the variable**

```go
// Reassigning a string variable
someCard = "10 of Diamond"
// Reassigning an int variable
someInt = 2000
```

## Functions

- In Go, there are 2 principal types of functions:

***`main()` Function***|***Helper Functions***
:-|:-
Only one per project|Can be multiple per project
Contained in the `main.go` file|Contained in differently-named `.go` files
Declared with `package main`|Declared with different package names
This is the entry-point of execution of an executable|These are re-usable blocks of logic

### Basic Function Declaration Format

```go
// func <name>(<args?>) <returnType?> { <body> ... <return?> }
func newCard() string {
    return "5 of Diamonds"
}
```

- `func`
  - Keyword to declare a function
- `name`
  - Name of the function
- `args`
  - Arguments of the function
  - Arguments can be multiple
  - Arguments are optional
  - Argument types must be specified
- `returnType`
  - The type of the value returned by the function
  - **A function returning a value needs to explicitly declare its `returnType` in its declaration**
  - **If the function returns nothing (e.g. only prints to the screen), skip the `returnType`**
- `body`
  - The body of the function's logic
  - Typically ends with a `return` statement to return the value from the function
  - However, `return` is optional and can be skipped (Match the `returnType` of the function)
  - A function returning a value needs to declare its `returnType` in its declaration

```go
// Package
// *******
package main

// Imports
// *******
import "fmt"

// Helper Functions
// ****************
func newCard() string {
    return "5 of Diamonds"
}

func getAge() int {
    return 20
}

// Main Function
// *************
func main() {
    // We are calling a function and assigning its return value to the variable
    // When calling a function, the return type of the function becomes the type of the variable it is assigned to
    var card string = newCard() // string
    var age int = getAge()      // int

    // Making use of the variables
    // All declared variables must be used
    fmt.Println(card)
    fmt.Println(age)
}
```

### Tuple-Like Assignement And Usage

- We can return multiple values using tuple-like
  - **On function, make sure to annotate the `returnType` using a tuple-like format**
- We can also assign multiple variables using tuple-like unpacking

```go
// A function that return a tuple-like (deck, deck)
func deal(d deck, handSize int) (deck, deck) {
    // Split the original deck into 2 using the handSize
    hand := d[:handSize]
    remainingDeck := d[handSize:]

    // Return the "hand" and the "remaining deck" as a tuple
    return hand, remaining_deck
}
```

## Arrays & Slices

2 types of data structures in Go for handling lists of records

### Array

- Basic list of values
- 0-based index
- Same element access syntax as typical lists and arrays
- **Fixed-length**
- Primitive Data Structure for lists
- **All of its elements must have the same type**
- Useful when needing a static list of constants

```go
// Array declaration format
variable := [length]type{csvValues}

// Example of Array
days := [7]string{
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
    "Sunday"
}
```

#### Accessing Array Element

- Same element access syntax as typical lists and arrays

```go
// Accessing an array element
today := days[0]
fmt.Println("Today is", today)
```

### Slice

- A bit advanced list of values
- 0-based index
- Same element access syntax as typical lists and arrays
- **Flexible-length: Can grow or shrink in length**
- **All of its elements must have the same type**
- Useful when needing to work on dynamic lists

```go
// Slice declaration format
variable := []type{values}

// Example of a Slice
cards := []string{
    "Ace of Diamond",
    newCard(),
    newCard()
}

// Example of a Slice
fruits := []string{
    "apple",
    "banana",
    "grape",
    "orange",
}
```

#### Accessing Slice Element

- Same element access syntax as typical lists and arrays

```go
// Accessing a slice element
fruit := fruits[-1]
fmt.Println("My fruit is", fruit)
```

#### Selecting A Subset/Range Of Slice

- This also follows the typical pattern of slices in other languages
- Also, the *up-to-index* is *up-to-but-not-including*

```go
// slice[startIndex: upToIndex]
twoFruits := fruits[0:2]

fmt.Println("two_fruits:")
for _, fruit := range twoFruits {
    fmt.Println("-", fruit)
}
```

- We can also use inference for the beginning or end of slice
  - If we skip the `startIndex`, we grab everything before the `upToIndex`

```go
// Skipping the start_index
allFruitsButLast := fruits[:len(fruits)-1]

fmt.Println("allFruitsButLast:")
for _, fruit := range allFruitsButLast {
    fmt.Println(fruit)
}
```

- If we skip the `upToIndex`, we grab everything after the `startIndex`

```go
// Skipping the upToIndex
allFruitsButFirst := fruits[1:]

fmt.Println("allFruitsButFirst:")
for _, fruit := range allFruitsButFirst {
    fmt.Println("-", fruit)
}
```

- If we skip both, we grab everything

```go
// Skipping both startIndex and upToIndex
allFruits := fruits[:]

fmt.Println("allFruits:")
for _, fruit := range allFruits {
    fmt.Println(fruit)
}
```

#### Appending New Elements To A Slice

- Because Slice is dynamic in length, we can add new elements to it
- **`append()` is a Pure Function**
  - **Appending does not modify the existing value**
  - **Instead, it returns a new value with the modification added**
  - **We have to set it back to the original variable**

```go
// Appending a new
cards = append(cards, "6 of Spades")
```

### Iteration With `for`-Loops

- We can iterate over both arrays or slices

#### Iteration Over Finite Sets

- `for`-loops are typically for iterating over a closed-set (finite set) of elements

```go
for index, card := range cards {
    fmt.Println(index, "--", card)
}
```

- `range <slice>`
  - The range of slice we want to iterate over
- `:=`
  - *With `for`-loops, the iteration variables are re-declared at each iteration*
  - So we have to use `:=` instead of `=`
- `index, card`
  - Variables used within the `for`-loop block
  - **Every declared variable must be used**
  - If either `index` or `card` is not going to be used in the loop body, replace with `_`

```go
cards := []string{
    newCard(),
    newCard(),
    newCard()
}

fmt.Println("Using for-loop with range:")
for index, card := range cards {
    fmt.Println(index, "--", card)
}

fmt.Println("Using for-loop with range without using index:")
for _, card := range cards {
    fmt.Println("--", card)
}
```

#### Iteration Over Infinite Sets

- In Go, there is no `while` keyword for doing iterations over infinite sets
- Instead, `for` can also be used in a `while`-like style for iterating over infinite sets of elements

```go
cards := []string{
    newCard(),
    newCard(),
    newCard()
}

fmt.Println("Using for-loop in a While-like style:")

i := 0
for i < len(cards) {
    fmt.Println("--", cards[i])
    i = i + 1
}
```

#### Iteration in C-style

- `range` is typically used with Go's `for`-loop
- However, we can always fallback to a C-style of `for` as well

```go
cards := []string{
    newCard(),
    newCard(),
    newCard()
}

fmt.Println("Using for-loop in a C-like-for-loop style:")
for i := 0; i < len(cards); i++ {
    fmt.Println("--", cards[i])
}
```

#### `break`, `continue`, And `do-while`

- Go does not have a `do-while` loop either
- Similar in other programming language, we can use `break` and `continue` to manipulate the flow of the loop
- We can also get an infinite loop if we use `for` without any conditions
  - Using this and `break`, we can get a `do-while`-like loop using `for`

```go
cards := []string{
    newCard(),
    newCard(),
    newCard()
}

fmt.Println("Using for-loop in an infinite-loop with break (Do-While-like) style:")

j := 0
for {
    // Do something at least once
    fmt.Println("--", cards[j])
    j += 1
    // Then check the condition:
    // Make sure it is reachable to avoid an infinite loop
    if j >= len(cards) {
        break
    }
}
```

## Types & Receiver functions

### Types

- **Go is not an *Object-Oriented* language**
- It does not have any comprehension of *Object* and *Class* types
- Instead, we use *Types* and *Receivers (Methods)*
  - Abstracted primitive types with additional functionalities
  - We want to *extend* a base type and add some extra functionalities to it
  - We could think of *Type* as a very simplified version of a *Class*

```go
// Declaring a type:
// type <typeName> <equivalentType>
type deck []string
```

- `type`
  - Keyword to declare a new type
- `typeName`
  - The name of the type
- `equivalentType`
  - The primitive type that is equivalent to the declared type

### Initializer Function

- Because Go is not an OOP language, it does not have a *Constructor* for the types
- Instead, we used an *Initializer* function that acts as a type-instance generator function

```go
// Initializes and returns a new deck of cards.
func newDeck() deck {
    // A deck is just an abstraction of a slice of strings
    cards := deck{}

    // Suits: An array of strings
    suits := [4]string{"Spade", "Diamond", "Heart", "Club"}

    // Values: An array of strings
    values := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

    // Build the combinations of Suits and Values
    for _, suit := range suits {
        for _, value := range values {
            // Create the new card
            newCard := fmt.Sprintf("%s of %s", value, suit)
            // Append the new card to the deck
            cards = append(cards, newCard)
        }
    }

    // Return the new deck
    return cards
}
```

### Receiver Function

- Receiver functions are like *Methods* that we attach to *Types*
- Receiver functions are called like *Methods* on type instances
- When attaching to a type, we typically use the initial of the type as the `this` or `self` keywords within the function to refer to the *Instance* of the type
  - This is not a mandate but a generally-accepted convention
  - Example: `deck` -> `d`

#### General Declaration Format Of A Receiver Function

```go
func (<t> <type>) <funcName>(<args>) <returnType> {
    <body>
}
```

- `<type>`
  - The type that we are attaching the receiver function to
- `<t>`
  - The *Instance Variable*
  - With Go, we never use `this` or `self`
  - Instead, by convention, we typically use the initial of the type
  - **NOTE: When the instance variable is not being used in the function, we can remove it**

#### Example Of A Receiver Function

```go
// Declaring a Receiver Function: Attaching to a deck type
// Returns a tuple-like
func (d deck) deal(handSize int) (deck, deck) {
    // Split the original deck into 2 using the handSize
    hand := d[:handSize]
    remDeck := d[handSize:]

    // Return the "hand" and the "remaining deck"
    return hand, remDeck
}
```

- When using a receiver function, we generally use it like a *Method* on the *Instance* of the type
- **NOTE: We can return multiple values from a Go function**
  - If we do not need one of the returned values, we can ignore by assigning it to `_`

```go
cardDeck := newDeck()
hand, _ := cardDeck.deal()
```

### Using Types

- Typically, types and their functionalities would be defined in a separate `.go` file
  - Using the same `package` to link them all inside the same project
- After declaring types and their functionalities *Receiver functions*, we can make use of them in our executable `main`

```go
// This is the main entry of the application
func main() {
    // Declaring and Initializing variable deck type
    // playingDeck is essentially a slice of strings
    playingDeck := newDeck()

    // Calling Type Receiver Function: Deal 5 cards
    hand, playingDeck := playingDeck.deal(5)

    // Calling Type Receiver Function: Print to screen
    hand.print()
}
```

## Writing To File

- To deal with underlying Operating System files such as text files, we make use of the `os` standard package
- Use `os.WriteFile()` to write to a system file

```go
import "os"

os.WriteFile(
    filename string,
    data []byte,
    permissions FileMode
)
```

- `filename`
  - A path of the file to write
- `[]byte`
  - Essentially a string of characters in binary format
  - Every element inside a *Byte Slice* correspond to an ASCII character code
  - We can use [asciitable.com](https://asciitable.com) as Xwalk table for the `Dec` column for better comprehension
  - Essentially, a *Byte Slice* is just another way to represent a string
- `permissions`
  - Unix-like permission in Octal format
  - Returns an `error` type by default

```go
import (
    "fmt"
    "os"
)
```

- Because `WriteFile()` only takes *Byte Slice*, we need to convert the string to write to file into a *Byte SLice* `[]byte`

```go
// Converting a string to a []byte
greetingStr := "Hello World!"
greetingBytes := []byte(greetingStr)

fmt.Println(greetingBytes)
```

- Now, we can use `WriteFile()`
  - We will just use a permission of `0o666` (read/write)
  - `WriteToFile()` returns an error type by default
    - We can use that with error handling later
    - For now, we will just ignore it

```go
// Writing text to file and ignoring returned error type
_ = os.WriteFile(
    "datasave_hello_world.tmp",
    greetingByte,
    0o666
)
```

## Reading From File

```go
import "os"

os.ReadFile(filename string)
```

- To deal with underlying OS files such as text files, we make use of the `"os"` standard package
- Use `ReadFile()` to read from a file
- `filename`
  - A path of the file to read from
  - Returns a *Byte-Slice*
  - Returns an `error` type by default

```go
import (
    "fmt"
    "os"
)
```

- Reading from the previously stored file

```go
// Reading from a file
greetingByte, err := os.ReadFile(
    "datasave_hello_world.sav"
)

// Converting Byte-Slice back to string
fmt.println(string(greetingByte))
```

### Error Handling

- Go does not have a `try-catch` similar-clause
- Instead, it returns errors as part of the function call
- If there was any error in reading the file, returned `err` will be *not `nil`*

```go
// Function call: Error would be returned as a second argument
greetingByte, err := os.ReadFile("datasave_hello_world.sav")

// Error Handling
if err != nil {
    // We have an error: Handle to resolution
    fmt.Println("Error:", err, "Creating a new file now.")
    greetingStr := "Hello World!"
    greetingByte = []byte(greetingStr)
}

// If here, then there was no error
fmt.Println(greetingByte)
```

- Instead of handling the error, we could also stop the execution completely

```go
// Function call: Error would be returned as a second argument
greetingByte, err := os.ReadFile("datasave_hello_world.sav")

// Error Handling
if err != nil {
    // We have an error: Stop execution and panic
    panic(err)
    os.Exit(1) // 0 is success, anything else is fail
}

// If here, then there was no error
fmt.Println(greetingByte)
```

## Unit Test

### How Do We Know What To Test?

- What makes sense
- What do you really care about with the feature?

### Testing In Go

- Go does not have a very strong unit test framework
- Very small set of functions for testing
  - Not similar to using typical Testing Framework
  - **We write Go codes to test Go codes**
- Create a new file ending in `_test.go`
  - E.g. For testing `deck.go`, use `deck_test.go`
- Define the test functions with `Test_` prefix
  - These `Test_` functions will be automatically called with `t *testing.T`
  - **The name after the `Test_` prefix does not necessarily need to match an existing function name**
  - `t` is the test-handler
    - If something is wrong, we use `t` to notify with an error message
    - `t.Errorf()`
      - Allows to return an error with string formatting
- To run all the tests in the package: `> go test`
  - **Make sure to run this from the location where the test files are located**

```go
// In deck_test.go
func Test_newDeck(t *testing.T) {
    // Deck Instance to test on
    var d deck

    // TEST CASE 1: A deck should be created with x number of cards
    // ------------------------------------------------------------
    // Create a new deck
    d = newDeck()

    // Set expectations:
    // The deck should have 52 number of cards
    expectedDeckLen := 52
    actualDeckLen := len(d)

    // Test expectations
    if expectedDeckLen != actualDeckLen {
        // If not, something is wrong --> Notify the test-handler t
        t.Errorf("Test Case 1: Expected deck length of %v. Got %v", expectedDeckLen, actualDeckLen)
    }
}
```

- When testing with files, we have to make sure that we cleanup the files we test with
  - **Go does not automatically take care of cleaning up test files**

```go
// In deck_test.go
func Test_saveToFileAndNewDeckFromFile(t *testing.T) {
    // Delete any file _decktesting.tmp from past tests if any
    os.Remove("_decktesting.tmp")

    // Create a new deck
    d := newDeck()

    // Save the deck to file
    d.saveToFile("_decktesting.tmp")

    // Attempt to load from disk
    loadedDeck := newDeckFromFile("_decktesting.tmp")

    // Set expectations:
    // The length of the loaded deck from file should be the same as the original deck
    expectedLoadedDeckLen := len(d)
    actualLoadedDeckLen := len(loadedDeck)

    // Test expectations
    if expectedLoadedDeckLen != actualLoadedDeckLen {
        // If not, something is wrong --> Notify the test-handler t
        // Errorf() is a formatted string: We can use % for placeholders
        t.Errorf("Expected loaded deck length of %v. Got %v", expectedLoadedDeckLen, actualLoadedDeckLen)
    }

    // Finally, clean up any temp test files
    os.Remove("_decktesting.tmp")
}
```
