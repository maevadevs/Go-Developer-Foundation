# Hello World In Go

## How is a Go file organized?

In practice, always the same typical pattern

1. **Package**
2. **Imports**
3. **Functions**

## Package

- Package == Project == Workspace == Namespace
- A Package can have many files associated with it
  - Each file ending with `.go`
  - The very first line of each file in the same package must start with the package declaration
- There are 2 types of packages:

1. **Executable**
     - Always `main` package
     - Must always have a function called `main()` as well
     - Generates a .exe/.bin file that we can run after compiling
     - Code used when we want to do something (Executable codes)
2. **Reusable**
     - Any other package name other than `main`
     - Code used as *helpers* for reusable logic, libraries, dependencies

- **Files in the same package do not have to be imported into each other before they can be used**
  - If other packages are declared with `package main`, the functions they contains can be used here directly
  - But one of the file must contain the `main()` function as the primary entry point of execution

```go
package main
```

## Import

- Allows to import reusable codes from other packages
- `"fmt"`
  - A standard package within Go ([pkg.go.dev/std](https://pkg.go.dev/std))
  - Short for *format*
  - Mostly used for debugging and development
- Unless we import a package, we have no access to any functionalities from another package
- We are not limited to importing packages from the standard library
  - We can import packages written by other engineers as well

```go
import "fmt"
```

## Function

- A `main` package must have a `main()` function as the entry-point
- This is a Go function, similar functionality to functions in other languages

```go
func main() {
    fmt.Println("Hello World!")
}
```

## How do we run code in our project?

```sh
> go run <filename>
> go run "<filepath>"
```

## Available Go commands

Go Commands | Action
:--|:--
`go build`|Compiles a bunch of Go source code files into executable binaries
`go run`|Compiles and execute a bunch of Go source code files (build + run) but does not produce an actual executable
`go fmt`|Formats all the code in each file in the current directory
`go install`|Compiles and *installs* a package
`go get`|Download the raw source code of someone else's package
`go test`|Runs any tests associated with the current projects
