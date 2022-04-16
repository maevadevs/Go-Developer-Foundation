// Package
// *******
package main

// Import
// ******
import "fmt"

// Type Definitions
// ****************
type hashmap map[string]string

// Functions
// *********
func main() {

	// A map of <string> => <string>
	base_colors := map[string]string{
		"red":   "ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
	}

	// Print the map to see its format
	fmt.Println("base_colors:", base_colors)

	// Declare a map with make()
	colors := make(map[string]string)
	fmt.Println("colors:", colors)

	// Add values to a map
	colors["white"] = "ffffff"
	colors["black"] = "000000"
	fmt.Println("colors:", colors)

	// Delete an existing mapping inside a map
	delete(colors, "white")
	fmt.Println("colors:", colors)

	// Defining colors as a hashmap type
	b_colors := hashmap{
		"red":   "ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
		"white": "ffffff",
		"black": "000000",
	}

	// Print colors using iteration
	b_colors.print()

}

// Receiver functions for hashmap
// ******************************
func (c hashmap) print() {

	// Iterating over a map of colors
	for color, hex := range c {
		fmt.Printf("Colors{%s: %s}\n", color, hex)
	}

}

// FOR WINDOWS: FROM PROJECT FOLDER
// 	To run:					go run src\main.go
// 	To compile:				go build -o bin\program.exe src\main.go
// 	To run after compile:	.\bin\program.exe

// FOR LINUX: FROM PROJECT FOLDER
// 	To run:					go run src/main.go
// 	To compile:				go build -o bin/program src/main.go
// 	To run after compile:	./bin/program
//	Compile + Run:			go build -o bin/program src/main.go && ./bin/program
