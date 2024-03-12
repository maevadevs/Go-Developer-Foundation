# Hello World In Go

---

- [How Is A Go File Organized?](#how-is-a-go-file-organized)
  - [Package](#package)
  - [Import](#import)
  - [Function](#function)
- [How Do We Run Code In Our Project?](#how-do-we-run-code-in-our-project)
- [Available Go Commands](#available-go-commands)
- [Go Modules](#go-modules)

---

## How Is A Go File Organized?

In practice, always the same typical pattern

1. **Package Declaration Declaration**
2. **Imports Packages Packages**
3. **Define Define Functions & Objects & Objects**

### Package

- Package == Project == Workspace == Namespace
- **A collection of source files `.go` in the same directory that are compiled together**
  - *Objects defined in one source file are visible to all other source files within the same package*
- A Package can have many files associated with it
  - Each file ending with `.go`
  - The very first line of each file in the same package must start with the package declaration
  - This is how we tie all files that are of the same project/package
  - *Package* is more of an abstract way to organize `.go` files
- There are 2 types of packages:

1. **Executable**
     - ***Always `main` package***
     - ***Must always have a function called `main()` as well***
     - Generates a `.exe`/`.bin` file that we can run after compiling
     - Code used when we want to *do* something: *Executable codes*
2. **Reusable**
     - ***Any other package name other than `main`***
     - Code used as *helpers* for reusable logic, libraries, dependencies

- *Files in the same package do not have to be imported into each other before they can be used*
  - If other files are declared with `package main`, the functions they contain can be used here directly
  - *But one of the file must contain the `main()` function as the primary entry point of execution*

```go
package main
```

### Import

- Allows to import reusable codes from other packages
- ***`import` is not necessary if the function/file to import is declared with the same package***
- Example: `"fmt"`
  - A standard package within Go ([pkg.go.dev/std/fmt](https://pkg.go.dev/fmt))
  - Short for *format*
  - Mostly used for debugging and development
- Unless we import a package, we have no access to any functionalities from another package
  - Exception: ***Files in the same package do not have to be imported into each other before they can be used***
- We are not limited to importing packages from the standard library
  - We can import packages written by other engineers and 3rd partie as well

```go
import "fmt"
```

### Function

- This is a Go function, similar functionality to functions in other languages

```go
// func funcName(arg argType) OptionalReturnType { body }
func myFunc(name str, age int) str {
    return fmt.Sprintf("Hello! My name is %s and I am %d years old!", name, age)
}
```

- A `main` package must have a `main()` function as the entry-point of the execution

```go
func main() {
    fmt.Println("Hello World!")
}
```

## How Do We Run Code In Our Project?

- Run this command in the console

```sh
# If current working directory is the project directory
$ go run <filename>

# For executing from anywhere
$ go run "/full/path/to/file.go"
```

## Available Go Commands

Command | Action
:-|:-
`go build`|**Compiles** a bunch of Go source code files into executable binaries
`go run`|**Compiles and executes** a bunch of Go source code files (build + run) *but does not produce an actual executable*
`go fmt`|When run in a directory with `.go` files, **formats** all the code in each file in the directory
`go install`|**Compiles** and **installs** a package
`go get`|**Download** the raw source code of someone else's package
`go test`|Runs any **tests** associated with the current projects

## Go Modules

```txt
Module > Package > Source Files
```

- *Module* contain one or more *Packages*
- *Package* is made of source files `.go`
  - **A collection of source files `.go` in the same directory that are compiled together**
  - Package is more of an abstract way of organizing and grouping `.go` files
  - But is best to organize them into a concrete directory structure: 1 Package = 1 Directory
  - ***The directory the code resides in is a package name by default***
  - ***Files in the same package do not have to be imported into each other before they can be used***
  - ***Files from different packages must be imported first before usage***

```tree
ExeModule/
|- go.mod
|- go.sum
|- src/
   |- main.go  -- package main
   |- db.go    -- package helpers
   |- files.go -- package helpers

LibModule/
|- go.mod
|- go.sum
|- src/
   |- sqldb/
   |  |- mysql.go   -- package sqldb
   |  |- postgre.go -- package sqldb
   |- files/
   |  |- pdf.go     -- package files
   |  |- csv.go     -- package files
   |- webserver/
      |- http.go    -- package webserver
      |- sftp.go    -- package webserver
```

- Initialize a directory into a Module with `go mod init path/for/module`

```sh
cd path/to/module/directory
go mod init unique/path/for/source
```

- This creates a `go.mod` file inside of the module folder
  - Contains the unique path for the module
  - Contains the version of go used for the module
- Typically, the path is on Github but it does not have to: It can be any path
