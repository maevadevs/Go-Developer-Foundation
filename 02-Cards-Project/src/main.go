// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Project Structure
// *****************
// 02-Cards-Project (Module)
// |- main.go      - Executable
// |- deck.go      - Describes what a Deck type is and how it works
// |- deck_test.go - Automated tests for deck.go

// Functions
// *********

// This is the main entry of the application.
func main() {
	// Variables: deck type
	// playingDeck is essentially a slice of strings
	playingDeck := newDeck()

	// Calling Type Receiver Function: Deal 5 cards
	hand, playingDeck := playingDeck.deal(5)

	// Calling Type Receiver Function: Print to screen
	fmt.Print("Current Hand: ")
	hand.print()
	fmt.Println("---")

	// Convert deck to string an print
	fmt.Print("Remaining Playing Deck: ")
	fmt.Println(playingDeck.toString())
	fmt.Println("---")

	// Save the playingDeck to file
	// 0. Save currentPath
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 1. Get the Project Path via go command
	goCmd := exec.Command("go", "list", "-m", "-f", "'{{.Dir}}'", "github.com/maevadevs/Go-Developer-Foundation/Cards-Project")
	projectPathBytes, err := goCmd.Output()
	if err != nil {
		panic(err)
	}

	// 2. Clean the project path: Convert to string, Remove any \n and '
	projectPath := strings.Trim(strings.Trim(string(projectPathBytes), "\n"), "'")

	// 3. Change the execution context into the project directory
	if err := os.Chdir(projectPath); err != nil {
		panic(err)
	}

	// 4. Check if the sav directory exists: If not, create it
	savPath := filepath.Join(projectPath, "sav")
	fmt.Printf("--- Saving playingDeck to file in %s/datasave_current_deck.sav\n", savPath)

	if _, err := os.Stat(savPath); err != nil {
		if os.IsNotExist(err) {
			// Directory does not exist: create new
			fmt.Printf("%s does not exist: Creating...\n", savPath)
			os.Mkdir(savPath, os.ModePerm)
		}
	}

	fmt.Println()

	// 5. Return to the original execution context
	err = os.Chdir(currentPath)
	if err != nil {
		panic(err)
	}

	// 6. Save the file in the savPath
	if err := playingDeck.saveToFile(fmt.Sprintf("%s/datasave_current_deck.sav", savPath)); err != nil {
		panic(err)
	}

	// Testing reading from the saved file
	fmt.Println("--- Reading playingDeck from saved file --- ")
	playingDeck = newDeckFromFile(fmt.Sprintf("%s/datasave_current_deck.sav", savPath))
	fmt.Println(playingDeck.toString())

	// Testing Shuffling
	playingDeck = newDeck()
	fmt.Println("---")
	fmt.Println("Before Shuffling The Deck:")
	fmt.Println(playingDeck.toString())
	fmt.Println("---")
	fmt.Println("After Shuffling The Deck:")
	playingDeck.shuffle(5)
	fmt.Println(playingDeck.toString())
}

// FOR WINDOWS:
//  To run:                 go run 02-Cards-Project\src\*.go
//  To compile:             go build -o 02-Cards-Project\bin\Program.exe 02-Cards-Project\src\*.go
//  To run after compile:   .\02-Cards-Project\bin\Program.exe
//  Compile + Run:          go build -o 02-Cards-Project\bin\Program.exe 02-Cards-Project\src\*.go && .\02-Cards-Project\bin\Program.exe

// FOR LINUX:
//  To run:                 go run 02-Cards-Project/src/*.go
//  To compile:             go build -o 02-Cards-Project/bin/Program 02-Cards-Project/src/*.go
//  To run after compile:   ./02-Cards-Project/bin/Program
//  Compile + Run:          go build -o 02-Cards-Project/bin/Program 02-Cards-Project/src/*.go && ./02-Cards-Project/bin/Program
