package main

import "fmt"

// NOTE: GO IS NOT OOP
// *******************
// There is no use of Classes inside of Go
// Instead, we use "types" and "Receivers"
// Go has its basic data types:
//	- We want to "extend" a base type and add some extra functionalities to it
//	- To work with them, we will use "Functions As A Receiver"
//		- Similar to method that can work with specific custom types

// Project Structure
// *****************
// 02-Cards
//	|- main.go - Executable
//	|- deck.go - Describes what a Deck type is and how it works
//	|- deck_test.go - Automated tests for deck.go

// When executing a programm with multiple files: > go run main.go deck.go

func main() {

	// Variables: deck type
	playing_deck := newDeck() // deck is a slice of strings

	// Calling Type Receiver Function: Deal 3 cards
	hand, playing_deck := playing_deck.deal(5)

	// Calling Type Receiver Function: Print to screen
	hand.print()
	fmt.Println("---")

	// Convert deck to string an print
	fmt.Println(playing_deck.toString())
	fmt.Println("---")

	// Save the deck to file
	playing_deck.saveToFile("sav_files/datasave_current_deck.sav")
	playing_deck = newDeck()

	// Testing reading from the saved file
	fmt.Println("--- Reading from saved file --- ")
	playing_deck = newDeckFromFile("sav_files/datasave_current_deck.sav")
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

// FOR WINDOWS: FROM PROJECT FOLDER
// 	To run:					go run src\main.go src\deck.go
// 	To compile:				go build -o bin\program.exe src\main.go src\deck.go
// 	To run after compile:	.\bin\program.exe

// FOR LINUX: FROM PROJECT FOLDER
// 	To run:					go run src/main.go src/deck.go
// 	To compile:				go build -o bin/program src/main.go src/deck.go
// 	To run after compile:	./bin/program
//	Compile + Run:			go build -o bin/program src/main.go src/deck.go && ./bin/program
