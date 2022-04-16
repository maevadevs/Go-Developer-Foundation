# Learning Go Through `Cards` Project

In this review:

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

- Variables can be initialized outside of a function
- But can only be assigned a value inside a function

### Basic variable declaration format

```go
// var <name> <type> = <value>
var my_card string = "Ace of Spade"
```

- `var`
  - Keyword to create a new variable
  - *Every declared variables must be used*
- `my_card`
  - Name of the variable
- `string`
  - Data type of the variable
  - Go is a statically-typed language
  - The variable type follows the variable name
- Go fundamental types:
  - `string`
  - `bool`
  - `int`
  - `float64`

- We can also declare only, then assigned a value later
  - When using this, Go assign the *null*-equivalent default value of the type to the declared variable

```go
var some_card string        // Declare variable: default value = ""
var some_int int            // Declare variable: default value = 0

some_card = "5 of Heart"    // Assign value to variable
some_int = 1001             // Assign value to variable
```

### Shortcut variable declaration format

- Go can also automatically infer the variable type from the assigned value
  - We use `:=` and omit the `var` keyword
  - **Only use `:=` when declaring a new variable *WITH* initialization *AND* type inference**

```go
// These are equivalent to the above declarations
some_card := "5 of Heart"
some_int := 1001
```

### Re-assigning values to variables

- Obviously, we can re-assign value to any variable
- **However, make sure to use `=` instead of `:=` when re-assigning**
  - `:=` is only used for initializing with type inference
  - `=` is used for all successive re-assignments
- **Make sure that the type of the re-assigned value matches the declared type of the variable**

```go
some_card = "10 of Diamond"
```

## Functions

- In Go, there are 2 principal types of functions:
  - `main()` function
    - Only one per project
    - Contained in the `main.go` file
    - Declared with `package main`
    - This is the entry-point of execution of an executable
  - Helper Functions
    - Can be multiple per project
    - Contained in differently-named `.go` files
    - Declared with different package names
    - These are re-usable blocks of logic

### Basic function declaration format

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
- `returnType`
  - The type of the value returned by the function
  - A function returning a value needs to declare its `returnType` in its declaration
  - If the function returns nothing (e.g. only prints to the screen), skip the `returnType`
- `body`
  - The body of the function's logic
  - Typically ends with a `return` statement to return the value from the function
  - However, `return` is optional and can be skipped (Matching the `returnType` of the function)
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
    card := newCard()   // string
    age := getAge()     // int

    // Making use of the variables
    fmt.Println(card)
    fmt.Println(age)

}
```

### Tuple-like assignement and usage

- We can return multiple values using tuple-like
  - On function, make sure to annotate the `returnType` using a tuple-like format
- We can also assign multiple variables using tuple-like unpacking

```go
// A function that return a tuple-like (deck, deck)
func deal(d deck, hand_size int) (deck, deck) {

    // Split the original deck into 2 using the hand_size
    hand := d[:hand_size]
    remaining_deck := d[hand_size:]

    // Return the "hand" and the "remaining deck" as a tuple
    return hand, remaining_deck

}
```

## Arrays & Slices

2 Types of Data Structures in Go for handling lists

### Array

- Basic list of values
- 0-based index
  - Same element access syntax as typical lists and arrays
- Fixed-length
- Primitive Data Structure for lists
- All of its elements must have the same type
- Useful when needing a static list of constants

```go
// Array declaration format
variable := [length]type{csv_values}

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

#### Accessing array element

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
- Flexible-length: Can grow or shrink in length
- All of its elements must have the same type
- Useful when needing to work on dynamic lists

```go
// Slice declaration format
variable := []type{values}

// Example of a Slice
cards := []string{
    newCard(),
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

#### Accessing slice element

- Same element access syntax as typical lists and arrays

```go
// Accessing a slice element
fruit := fruits[-1]
fmt.Println("Mys fruit is", fruit)
```

#### Selecting a subset/range of slice

- This also follows the typical pattern of slices in other languages
- Also, the *up-to-index* is *up-to-but-not-including*

```go
// slice[start_index: up_to_index]
two_fruits := fruits[0:2]

fmt.Println("two_fruits:")
for _, fruit := range two_fruits {
    fmt.Println("-", fruit)
}
```

- We can also use inference for the beginning or end of slice
  - If we skip the `start_index`, we grab everything before the `up_to_index`

```go
// Skipping the start_index
all_fruits_but_last := fruits[:len(fruits)-1]

fmt.Println("all_fruits_but_last:")
for _, fruit := range all_fruits_but_last {
    fmt.Println(fruit)
}
```

- If we skip the `up_to_index`, we grab everything after the `start_index`

```go
// Skipping the up_to_index
all_fruits_but_first := fruits[1:]

fmt.Println("all_fruits_but_first:")
for _, fruit := range all_fruits_but_first {
    fmt.Println("-", fruit)
}
```

- If we skip both, we grab everything

```go
// Skipping both start_index and up_to_index
all_fruits := fruits[:]

fmt.Println("all_fruits:")
for _, fruit := range all_fruits {
    fmt.Println(fruit)
}
```

#### Appending new elements to a slice

- Because Slice is dynamic in length, we can add new elements to it
- Appending does not modify the existing value
- Instead, it returns a new value with the modification added
- Pure Function: We have to set it back to the original variable

```go
// Appending a new 
cards = append(cards, "6 of Spades")
```

### Iteration with `for`-loops

- We can iterate over both arrays or slices

#### Iteration over finite sets

- `for`-loops are typically for iterating over a closed-set (finite set) of elements

```go
for index, card := range cards {
    fmt.Println(index, "--", card)
}
```

- `range <slice>`
  - The range of slice we want to iterate over
- `:=`
  - With `for`-loops, the iteration variables are re-declared at each iteration
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

#### Iteration over infinite sets

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

#### `break`, `continue`, and `do-while`

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
- Instead, we use *Types* and *Receivers*
  - Abstracted primitive types with additional functionalities
  - We want to *extend* a base type and add some extra functionalities to it
  - We could think of *Type* as a very simplified version of a *Class*

```go
// Declaring a type: type <typeName> <equivalentType>
type deck []string
```

- `type`
  - Keyword to declare a new type
- `typeName`
  - The name of the type
- `equivalentType`
  - The primitive type that is equivalent to the declared type

#### Initializing Function

- Because Go is not an Object-Oriented language, it does not have a *Constructor* for the types
- Instead, we used an initializer function that acts as a type-instance generator function

```go
func newDeck() deck {

    // A deck is just an abstraction of a slice of strings
    cards := deck{}

    // Suits: A slice of strings
    suits := []string{"Spade", "Diamond", "Heart", "Club"}

    // Values: A slice of strings
    values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

    // Build the combinations of Suits and Values
    for _, suit := range suits {
        for _, value := range values {
            // Create the new card
            new_card := value + " of " + suit
            // Append the new card to the deck
            cards = append(cards, new_card)
        }
    }

    // Return the new deck
    return cards

}
```

### Receiver function

- Receiver functions are like *Methods* that we attach to *Types*
- Receiver functions are called like *Methods* on type instances
- When attaching to a type, we typically use the initial of the type as the `this` or `self` keywords within the function to refer to the *Instance* of the type
  - This is not a mandate but a generally-accepted guideline
  - Example: `deck` -> `d`

#### General declaration format of a receiver function

```go
func (<t> <type>) <funcName>(<args>) <returnType> {
    <body> 
}
```

- `<type>` - The type that we are attaching the receiver function to
- `<t>` - The Instance Variable
  - With Go, we never use `this` or `self`
  - Instead, by convention, we typically use the initial of the type

#### Example of a receiver function

```go
// Declaring a Receiver Function: Attaching to a deck type
// Returns a tuple-like
func (deck d) deal(hand_size int) (deck, deck) {

    // Split the original deck into 2 using the hand_size
    hand := d[:hand_size]
    remaining_deck := d[hand_size:]

    // Return the "hand" and the "remaining deck" as a tuple
    return hand, remaining_deck

}
```

- When using a receiver function, we generally use it like a *Method* on the *Instance* of the type

### Using types

- Typically, types and their functionalities would be defined in a separate `.go` file
  - Using the same `package` to link them all inside the same project
- After declaring types and their functionalities (Receiver functions), we can make use of them in our executable (`main`)

```go
func main() {

    // Declaring and Initializing variable deck type
    // playing_deck is an abstraction of a slice of strings
    playing_deck := newDeck()

    // Using Receiver function: Deal 3 cards
    hand, playing_deck := playing_deck.deal(5)

    // Using Receiver function: Print the deck
    hand.print()

}
```

## Writing To File

- To deal with underlying OS files such as text files, we make use of the `"io/ioutil"` standard package
- Use `WriteFile(filename string, data []byte, permissions)` to write to a file
  - `filename` - A path of the file to write
  - `[]byte` - Essentially a string of characters
    - Every element inside a *Byte Slice* correspond to an ASCII character code
    - We can use [asciitable.com](https://asciitable.com) as XWalk table for comprehension
    - Essentially, a *Byte Slice* is just another way to represent a string
  - `permissions` - Unix-like permission in Octal format
  - Returns an `error` type by default

```go
import (
    "fmt"
    "io/ioutil"
)
```

- Because `WriteFile()` only takes *Byte Slice*, we need to convert the string to write to file into a *Byte SLice* `[]byte`

```go
// Converting a string to a []byte
greeting_str := "Hello World!"
greeting_bytes := []byte(greeting_str)

fmt.Println(greeting_bytes)
```

- Now, we can use `WriteFile()`
  - We will just use a permission of `0666` (read/write)
  - `WriteToFile()` returns an error type by default
    - We can use that with error handling later
    - For now, we will just ignore it

```go
// Writing text to file and ignoring returned error type
_ = ioutil.WriteFile(
    "datasave_hello_world.tmp", 
    greeting_byte, 
    0666
)
```

## Reading From File

- To deal with underlying OS files such as text files, we make use of the `"io/ioutil"` standard package
- Use `ReadFile(filename string)` to read from a file
  - `filename` - A path of the file to read from
  - Returns a *Byte-Slice*
  - Returns an `error` type by default

```go
import (
    "fmt"
    "io/ioutil"
)
```

- Reading from the previously stored file

```go
// Reading from a file
greeting_byte, err := ioutil.ReadFile(
    "datasave_hello_world.sav"
)

// Converting Byte-Slice back to string
fmt.println(string(greeting_byte))
```

### Error Handling

- Go does not have a `try-catch` similar-clause
- Instead, it returns errors as part of the function call
- If there was any error in reading the file, returned `err` will be *not `nil`*

```go
// Function call: Error would be returned as a second argument
greeting_byte, err := ioutil.ReadFile("datasave_hello_world.sav")

// Error Handling
if err != nil {
    // We have an error: Handle to resolution
    fmt.Println("Error:", err, "Creating a new file now.")
    greeting_str := "Hello World!"
    greeting_byte = []byte(greeting_str)
}

// If here, then there was no error
fmt.Println(greeting_byte)
```

## Unit Test

- Go does not have a very strong unit test framework
- Very small set of functions for testing
  - Not similar to using typical Testing Framework
  - We write Go codes to test Go codes
- Create a new file ending in `_test.go`
  - E.g. For testing `deck.go` --> `deck_test.go`
- Define the test functions with `Test` prefix
  - Functions starting With `Test` will be automatically called with `t *testing.T`
  - `t` is the test-handler
    - If something is wrong, we use `t` to notify with an error message
    - `t.Errorf()` - Allows to return an error with string formatting
- To run all the tests in the package: `> go test`

```go
// In deck_test.go

func Test_NewDeck(t *testing.T) {

    // CASE 1: A deck should be created with x number of cards
    // -------------------------------------------------------

    // Create a new deck
    d := newDeck()

    // Expect: The deck has 52 number of cards
    if len(d) != 52 {
        // If not, something is wrong --> Notify the test-handler t
        // Errorf() is a formatted string: We can use % for placeholders
        t.Errorf("Expected deck length of 52. Got %v", len(d))
    }

}
```

- When testing with files, we have to make sure that we cleanup the files we test with
  - Go does not automatically take care of test files

```go
// In deck_test.go

func Test_SaveToFileAndNewDeckFromFile(t *testing.T) {

    // Delete any file _decktesting.tmp from past tests if any
    os.Remove("_decktesting.tmp")

    // Create a new deck
    d := newDeck()

    // Save the deck to file
    d.saveToFile("_decktesting.tmp")

    // Attempt to load from disk
    loaded_deck := newDeckFromFile("_decktesting.tmp")

    // Expect: the Loaded Deck has 52 number of cards
    if len(loaded_deck) != 52 {
        // If not, something is wrong --> Notify the test-handler t
        // Errorf() is a formatted string: We can use % for placeholders
        t.Errorf("Expected loaded deck length of 52. Got %v", len(d))
    }

    // Finally, clean up after ourselves
    os.Remove("_decktesting.tmp")

}
```

### How do we know what to test?

- What makes sense
- What do you really care about with the feature?
