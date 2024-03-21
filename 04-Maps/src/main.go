// Package
// *******
package main

// Import
// ******
import "fmt"

// main
// ****

func main() {
	// A map of <string> => <string>
	baseColors := map[string]string{
		"red":   "ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
	}

	// Print the map to see its format
	fmt.Println("baseColors:", baseColors)
	fmt.Println()

	// Declare a map with make()
	colors := make(map[string]string)
	fmt.Println("colors:", colors)
	fmt.Println()

	// Add values to a map
	// We cannot use dot-syntax with maps
	colors["white"] = "ffffff"
	colors["black"] = "000000"
	fmt.Println("colors:", colors)

	// Delete an existing mapping inside a map
	delete(colors, "white")
	fmt.Println("colors:", colors)
	fmt.Println()

	// Defining colors as a custommap type
	b_colors := custommap{
		"red":   "ff0000",
		"green": "00ff00",
		"blue":  "0000ff",
		"white": "ffffff",
		"black": "000000",
	}

	// Print colors using iteration
	b_colors.print()
}

// FOR WINDOWS: FROM PROJECT FOLDER
//  To compile:             go build -o 04-Maps\bin\Program.exe .\04-Maps\src\
//  To run after compile:   .\04-Maps\bin\Program.exe

// FOR LINUX: FROM PROJECT FOLDER
//  To compile:             go build -o 04-Maps/bin/Program ./04-Maps/src/
//  To run after compile:   ./04-Maps/bin/Program
//  Compile + Run:          go build -o 04-Maps/bin/Program ./04-Maps/src/ && ./04-Maps/bin/Program
