/**
 * @file: Describes what a Deck type is and how it works
 */

// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Type Declaration
// ****************
// A Deck type is an abstraction of a slice of string with additional functionalities
type deck []string

// Initializer Function (Type Constructor)
// ***************************************
// newDeck()
// Initialize and create a new deck of cards
func newDeck() deck {

	// A deck is just an abstraction of a slice of strings
	cards := deck{}

	// Suits
	suits := []string{"Spade", "Diamond", "Heart", "Club"}

	// Values
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	// Build the combinations of Suits and Values
	for _, suit := range suits {
		for _, value := range values {
			new_card := value + " of " + suit
			cards = append(cards, new_card)
		}
	}

	// Return the new deck
	return cards

}

// Receiver Functions (Type Methods)
// *********************************
// deck.print()
// Receiver Functions for the deck type to print the value representation of a deck
func (d deck) print() {

	for _, card := range d {
		fmt.Println(card)
	}

}

// deck.Deal()
// Receiver Function to deal cards from the deck
func (d deck) deal(hand_size int) (deck, deck) {

	// Split the original deck into 2 using the hand_size
	hand := d[:hand_size]
	rem_deck := d[hand_size:]

	// Return the "hand" and the "remaining deck"
	return hand, rem_deck

}

// deck.ToString()
// Receiver Function to convert a deck into its string representation
func (d deck) toString() string {

	// deck -> []string
	d_strs := []string(d)

	// deck -> string: Condensce by joining with a separator
	// We can join a []string to string using Join(strs []string, sep string)
	d_str := strings.Join(d_strs, "|")

	// Finally return
	return d_str

}

// deck.Shuffle()
// Receiver Function that shuffle the deck
// Go does not have a standard way to randomize order in a slice
// So we will put our custom logic instead: With Time-Based Random Number Generator
func (d deck) shuffle() {

	// Time-Based Random Number Generator
	source := rand.NewSource(time.Now().UnixNano())
	rand_gen := rand.New(source)

	// Go through the list of cards
	for current_i := range d {

		// Generate a random index number: [0, len(d)-1]
		//	We can make use of the math.intn() function for this
		//	By default, the random number generator will always use the exact same seed
		//	Without a new seed, we will always get the exact same sequence
		//	We make use of the defined Time-Based Random Number Generator

		random_i := rand_gen.Intn(len(d) - 1)

		// Swap the current card and the card at the random index number
		// Same syntax as Python Tuple for swapping
		d[random_i], d[current_i] = d[current_i], d[random_i]
	}

}

// deck.SaveToFile()
// Receiver Function to save the deck to a file
// Using WriteFile(filename string, data []byte, permissions os.FileMode)
//	- Returns an error type if there is any
// 	- deck -> []string -> string -> []byte
func (d deck) saveToFile(filename string) error {

	// First, convert the deck to string: deck -> []string -> string
	d_str := d.toString()

	// Then convert this string to []byte to use with WriteFile()
	d_bytes := []byte(d_str)

	// Then write to file: Return its error if any
	return ioutil.WriteFile(filename, d_bytes, 0666)

}

// Regular Helper Functions
// ************************
// newDeckFromFile()
// Function to create a new deck from an existing save file
func newDeckFromFile(filename string) deck {

	// Read from the file
	deck_bytes, err := ioutil.ReadFile(filename)

	// Error Handling: Make sure that there is no errors before continuing
	if err != nil {
		// If here, we got an error in reading the file
		//	1. Print out the error
		//	2. Create a new deck
		fmt.Println("Error", err)
		fmt.Println("We are creating a new deck...")
		return newDeck()

		// Or another option would be to completely quit the program
		// We can do that with the "os" package
		// os.Exit(404)
	}

	// If here, there was no errors in reading the file
	// Convert deck_bytes back to an actual deck: []byte -> string -> []string -> deck
	deck_str := string(deck_bytes)

	// We get a psv (pipe-separated values)
	// We can use strings.Split(s string, sep string) to parse this into []string
	deck_strs := strings.Split(deck_str, "|")

	// We can use the slice of strings to convert into an actual deck
	return deck(deck_strs)

}
