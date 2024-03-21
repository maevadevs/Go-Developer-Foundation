// Package
// *******
package main

// Imports
// *******
import "fmt"

// Functions
// *********

// This is the main entry of the application.
func main() {
	// Print to the screen
	fmt.Println("Hello World!")

	// Print to screen with some string formatting
	// It does not insert new line by default
	// We can still insert new line explicitly
	fmt.Printf("Hello again %s", "World!\n")

	// Print to screen without new line
	// We can still insert new line explicitly
	fmt.Print("Hello with no new line!\n")

	// Wait before closing the console
	fmt.Println("\nPress enter to exit...")
	fmt.Scanln()
}

// FOR WINDOWS:
//  To run:                 go run 01-Hello-World\src\main.go
//  To compile:             go build -o 01-Hello-World\bin\Program.exe 01-Hello-World\src\main.go
//  To run after compile:   .\01-Hello-World\bin\Program.exe
//  Compile + Run:          go build -o 01-Hello-World\bin\Program.exe 01-Hello-World\src\*.go && .\01-Hello-World\bin\Program.exe

// FOR LINUX:
//  To run:                 go run 01-Hello-World/src/main.go
//  To compile:             go build -o 01-Hello-World/bin/Program 01-Hello-World/src/main.go
//  To run after compile:   ./01-Hello-World/bin/Program
//  Compile + Run:          go build -o 01-Hello-World/bin/Program 01-Hello-World/src/main.go && ./01-Hello-World/bin/Program
