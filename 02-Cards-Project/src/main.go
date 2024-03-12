// Package
// *******
package main

// Imports
// *******
import "fmt"

// Project Structure
// *****************

// 02-Cards
//	|- main.go - Executable
//	|- deck.go - Describes what a Deck type is and how it works
//	|- deck_test.go - Automated tests for deck.go

// Functions
// *********

func main() {
	// Variables: deck type
	playing_deck := newDeck() // deck is a slice of strings

	// Calling Type Receiver Function: Deal 5 cards
	hand, playing_deck := playing_deck.deal(5)

	// Calling Type Receiver Function: Print to screen
	hand.print()
	fmt.Println("---")

	// Convert deck to string an print
	fmt.Println(playing_deck.toString())
	fmt.Println("---")

	// Save the deck to file
	playing_deck.saveToFile("sav/datasave_current_deck.sav")

	// Testing reading from the saved file
	fmt.Println("--- Reading from saved file --- ")
	playing_deck = newDeckFromFile("sav/datasave_current_deck.sav")
	fmt.Println(playing_deck.toString())

	// Testing Shuffling
	playing_deck = newDeck()
	fmt.Println("---")
	fmt.Println("Before Shuffling The Deck:")
	fmt.Println(playing_deck.toString())
	fmt.Println("---")
	fmt.Println("After Shuffling The Deck:")
	playing_deck.shuffle()
	fmt.Println(playing_deck.toString())
}

// FOR WINDOWS:
//	To run:					go run 02-Cards-Project\src\*.go
//	To compile:				go build -o 02-Cards-Project\bin\Program.exe 02-Cards-Project\src\*.go
//	To run after compile:	.\02-Cards-Project\bin\Program.exe
//	Compile + Run:			go build -o 02-Cards-Project\bin\Program.exe 02-Cards-Project\src\*.go && .\02-Cards-Project\bin\Program.exe

// FOR LINUX:
//	To run:					go run 02-Cards-Project/src/*.go
//	To compile:				go build -o 02-Cards-Project/bin/Program 02-Cards-Project/src/*.go
//	To run after compile:	./02-Cards-Project/bin/Program
//	Compile + Run:			go build -o 02-Cards-Project/bin/Program 02-Cards-Project/src/*.go && ./02-Cards-Project/bin/Program
