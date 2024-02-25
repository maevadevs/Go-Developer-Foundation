# Interface

---

- [Interface Basics](#interface-basics)
- [Interface Rules](#interface-rules)
- [Exploring the `HTTP` Package](#exploring-the-http-package)
  - [Additional Interface Syntax](#additional-interface-syntax)
  - [The `io.Reader` Interface](#the-ioreader-interface)
  - [The `Read()` Function](#the-read-function)
  - [The `Writer` Interface](#the-writer-interface)
  - [The `io.Copy()` function](#the-iocopy-function)
  - [Implementation of `io.Copy()`](#implementation-of-iocopy)
  - [Custom `Writer`](#custom-writer)

---

## Interface Basics

- Interfaces are keys features in Go
- We know that:
  - Every value has a type
  - Every function has to specify the types of its arguments
  - So should every function accomodate different types even if the logic in it is identical?
- Let's imagine we have the following structs

```go
type englishBot struct{}
type spanishBot struct{}
```

- They are very similar to each other
- The following functions would probably be different for each bot
  - NOTE: For receiver, when the instance variable is not being used in the function, we can remove it

```go
func (englishBot) getGreeting() string {
    // Imagine some logic that is custom to englishBot
    // Custom implementation for englishBot
    return "Hello!"
}

func (spanishBot) getGreeting() string {
    // Imagine some logic that is custom to spanishBot
    // Custom implementation for spanishBot
    return "Hola!"
}
```

- However, the functionalities in the following functions are not specifically tied to a bot type
  - We could generalize them with any other types of *things* that are similar
  - In part, this is what Interfaces resolve: Makes it easier to re-use codes in our codebase

```go
func printGreeting(eb englishBot) {
    fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishhBot) {
    fmt.Println(sb.getGreeting())
}
```

- We can refactor those to make use of interface
  - Interface is like a contract with a type
  - As long as a `type` implements the same functions defined by the interface, they are good to go

```go
//  type <type_name> interface {
//      <func_name>(<list_arg_types>) (<list_return_types>)
//  }
type IBot interface {
    getGreeting() string
    getBotVersion() float64
    respondToUser(user) (string, error)
}
type englishBot struct{}
type spanishBot struct{}

func printGreeting(b IBot) {
    fmt.Println(b.getGreeting())
}
```

- Interface are essentially `type` definitions
- Interfaces define what functionalities a `type` should implement
  - It is a contract with a `type`
  - A `type` that wants to act as `IBot` must implement the functions defined by the `IBot` interface
  - Interfaces allows our code to be more DRY

## Interface Rules

- **Interfaces are NOT *Concrete Types***
  - **Concrete Type**
    - We can create a value from directly
    - Examples: `engishBot`, `spanishBot`, `int`, `string`...)
  - **Interface Type**
    - We cannot create values directly from an Interface (e.g. `IBot`)
    - Interface Types are only used to define arguments taken by functions
- **Interfaces are NOT *Generic Types***
  - Go does not have support for Generic Types
- **Interfaces are *Implicit***
  - As long as the concrete types follow the contract functions, they are a type of the interface
  - In Go, we do not use `IMPLEMENTS` keyword with interfaces
  - Might make it a bit confusing which types implement an interface
- **Interfaces are a *Contracts* to help us manage types**
  - But we stil need to know how to implement well the logic
  - If we don't, we will have *GIGO (Garbage-In, Garbage-Out)*

## Exploring the `HTTP` Package

- This is an example of using Interface in Go
- We will use the `net/http` package to make HTTP requests

```go
func main() {
    // Create an HTTP request
    resp, err := http.Get("https://example.com")

    // Error Handling
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    // Log out the response
    fmt.Println(resp)
}
```

- But the response we get is not actually the HTML representation
  - The response `resp` is a pointer
  - `resp` is actually a `struct` that contains information about the response object ([docs](https://pkg.go.dev/net/http@go1.18.1#Response))
  - `resp.Body` is of type `io.ReadCloser`
  - `io.ReadCloser` is an interface

```go
type ReadCloser interface {
    Reader
    Closer
}

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}
```

So overall, we have:

```sh
Response Struct {
    Status string
    StatusCode int
    Body ReadCloser --> ReadCloser interface {
        Reader --> Reader interface {
            Read(p []byte) (n int, err error)
        }
        Closer --> Closer interface {
            Close() error
        }
    }
}
```

### Additional Interface Syntax

- We can use interface as a type inside of a struct
  - It means that the field can have any type that can fulfill the contract of the specified interface
  - E.g. `ReadCloser` can take either a `Reader` or a `Closer`
- **In Go, we can take multiple interfaces and assemble them together to create another interface**
  - Meaning that *all the contracts of all specified interfaces must be satisfied*

```go
type ReadCloser interface {
    Reader  // This is an interface
    Closer  // This is an interface
}
```

### The `io.Reader` Interface

- One of the most common interfaces in Go
- It is possible for a program to read from different kinds of data sources (e.g. Text Files, HTTP Request Body, Image, User Inputs...)
  - Each one of those handler functions might return different data types
  - Each one of those handler functions might have different custom implementations
- Without interfaces, we would have to define handler functions for each types
  - Though those handler functions would have the same logic
  - The solution to make this more *DRY* is the `Reader` interface

```sh
[Different Src Types] --> [Reader] --> Universal Data Format
```

- We can think of the interface as an adapter to generalize
  - `Reader` requires to define a function that can output a `[]byte`
  - We can write any different types of functions that can do so

### The `Read()` Function

- **The `Reader` interface requires the implementation of a `Read()` function**
  - Called with `[]byte` and returns `(int, error)`
  - Takes the original raw data and feed it into the `[]byte`
  - `[]byte` is passed to `Read` from the thing that wants to consume the data
  - Remember that a `[]byte` is passed by reference
    - Modifying it means modifying the same object in memory
    - Takes advantage of the concept of pointers
  - `int` is the number of bytes that was read into that slice
  - `error` when something goes wrong

So our main function becomes like this:

```go
func main() {
    // Create an HTTP request
    resp, err := http.Get("https://example.com")

    // Error Handling
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    // A byte slice "pointer" for getting the http data
    // make(<type>, <number_of_elements>)
    bs := make([]byte, 99999)

    // Pass the byte slice to the read function
    resp.Body.Read(bs)

    // Print out the actual byte slice
    fmt.Println(string(bs))
}
```

### The `Writer` Interface

- Up to now, we have had the following diagram: Working with the `Reader` interface

```sh
[Different Src Types] --> [Reader] --> [Universal Data Format: []byte]
```

- Go has another interface that can do the exact opposite: `Writer` interface
  - Describes something that can take info and send it outside of the program
  - Requires to implement a `Write()` function

```sh
[Universal Data Format: []byte] --> [Writer] --> [Some form of Output]
```

- So we need to find something that *implements* the `Writer` interface
  - Use that to log out the data that we received from the `Reader`

### The `io.Copy()` function

- `io.Copy` implements the `Writer` interface
  - It requires something that implements the `Writer` interface
  - It requires something that implements the `Reader` interface
  - `Copy(Writer, Reader) (int64, error)`
- Here, we will only use the Standard Output (console) for now
  - `os.Stdout` implements the `Writer` interface
    - Actually of type `*File`, which implements the `Writer` interface
    - `*File` has a function `Write()`
  - `resp.Body` implements the `Reader` interface

```go
func main() {
    // Create an HTTP request
    resp, err := http.Get("https://example.com")

    // Error Handling
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    // Take the []byte from the resp.Body that is a Reader and output to the standard output
    io.Copy(os.Stdout, resp.Body)
}
```

### Implementation of `io.Copy()`

- Takes something that implements the `Writer` interface (i.e. has a `Write()` function) for writing into: `os.Stdout`
- Takes something that implements the `Reader` interface (i.e. has a `Read()` function) to read data from: `resp.Body`
- Pass them to `copyBuffer()`
  - This handles the creation of the byte slice and the internal logic of passing the byte slice to the `Read()` function

### Custom `Writer`

- Here is an example of a custom type that implements the `Writer` interface
  - **A `Writer` implementation must define a `Write()` function**

```go
type Writer interface {
    Write([]byte) (int, error)
}
```

- So we have to create a type that would fulfill this condition in order to create a `Writer`

```go
// Custom Interface
// ****************
type logWriter struct{}

// logWriter implements Writer
func (logWriter) Write(bs []byte) (int, error) {
    // Print the byte-slice
    fmt.Println(string(bs))
    // A custom implementation
    fmt.Print("Just wrote this many bytes: ", len(bs))
    // Return
    return len(bs), nil
}
```

Then, we can make use of that in `main()`

```go
func main() {
    // Create an HTTP request
    resp, err := http.Get("https://example.com")

    // Error Handling
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }

    // Using a custom type that implements the Writer interface
    lw := logWriter{}
    io.Copy(lw, resp.Body)
}
```
