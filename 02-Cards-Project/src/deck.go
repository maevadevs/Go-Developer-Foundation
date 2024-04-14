/**
 * @file: Describes what a Deck type is and how it works.
 */

// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Type Declaration
// ****************

// A Deck type is an abstraction of a slice of string with additional functionalities.
type deck []string

// Initializer Function (Type Constructor)
// ***************************************

// Initializes and returns a new deck of cards.
func newDeck() deck {
	// A deck is just an abstraction of a slice of strings
	cards := deck{}

	// Suits: An array of strings
	suits := [4]string{"Spade", "Diamond", "Heart", "Club"}

	// Values: An array of strings
	values := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	// Build the combinations of Suits and Values
	for _, suit := range suits {
		for _, value := range values {
			// Create the new card
			newCard := fmt.Sprintf("%s of %s", value, suit)
			// Append the new card to the deck
			cards = append(cards, newCard)
		}
	}

	// Return the new deck
	return cards
}

// Receiver Functions (Type Methods)
// *********************************

// deck.print()
// Receiver Functions for the deck type to print the value representation of a deck.
func (d deck) print() {
	dStrs := []string(d)
	dStr := strings.Join(dStrs, "|")
	fmt.Printf("%s", dStr)
	fmt.Println()
}

// deck.deal()
// Receiver Function to deal cards from the deck.
func (d deck) deal(handSize int) (deck, deck) {
	// Split the original deck into 2 using the handSize
	hand := d[:handSize]
	remDeck := d[handSize:]

	// Return the "hand" and the "remaining deck"
	return hand, remDeck
}

// deck.toString()
// Receiver Function to convert a deck into its string representation.
func (d deck) toString() string {
	// deck -> []string
	dStrs := []string(d)

	// deck -> []string -> string: Condensce by joining with a separator
	// We can join a []string to string using Join(strs []string, sep string)
	dStr := strings.Join(dStrs, "|")

	// Finally return
	return dStr
}

// deck.shuffle()
// Receiver Function that shuffle the deck.
// Go does not have a standard way to randomize order in a slice.
// So we use our custom logic instead: With Time-Based Random Number Generator.
func (d deck) shuffle(times uint) {
	// Time-Based Random Number Generator
	source := rand.NewSource(time.Now().UnixNano()) // seed
	randGen := rand.New(source)

	// Suffle whatever times was passed in: At least once
	if times == 0 {
		times = 1
	}
	for range times {
		// Go through the list of cards
		for currentI := range d {
			// Generate a random index number: [0, len(d)-1]
			// We could make use of the rand.intn() function for this
			// By default, the random number generator will always use the exact same seed
			// Without a new seed, we will always get the exact same sequence
			// randGen - We make use of the defined Time-Based Random Number Generator
			randomI := randGen.Intn(len(d) - 1)

			// Swap the current card and the card at the random index number
			// Same syntax as Python Tuple for swapping
			d[randomI], d[currentI] = d[currentI], d[randomI]
		}
	}
}

// deck.saveToFile()
// Receiver Function to save the deck to a file.
// Using os.WriteFile().
//   - Returns an error type if there is any
//   - deck -> []string -> string -> []byte
func (d deck) saveToFile(filename string) error {
	// First, convert the deck to string: deck -> []string -> string
	dStr := d.toString()

	// Then convert this string to []byte to use with WriteFile()
	dBytes := []byte(dStr)

	// Then write to file: Return its error if any
	return os.WriteFile(filename, dBytes, 0o666)
}

// Helper Functions
// ****************

// newDeckFromFile()
// Function to create a new deck from an existing save file
func newDeckFromFile(filename string) deck {
	// Read from the file
	deckBytes, err := os.ReadFile(filename)
	// Error Handling: Make sure that there is no errors before continuing
	if err != nil {
		// If here, we got an error in reading the file
		//   1. Print out the error but don't fail execution
		//   2. Create a new deck
		fmt.Println("Error:", err)
		fmt.Println("We are creating a brand new deck...")
		return newDeck()

		// Or another option would be to completely quit the program
		// We can do that with the "os" package
		// panic(err)
		// os.Exit(1) // 0 is success, anything else is fail
	}

	// If here, there was no errors in reading the file
	// Convert deckBytes back to an actual deck: []byte -> string -> []string -> deck
	deckStr := string(deckBytes)

	// We get a psv (pipe-separated values)
	// We can use strings.Split(s string, sep string) to parse this into []string
	deckStrs := strings.Split(deckStr, "|")

	// We can use the slice of strings to convert into an actual deck
	return deck(deckStrs)
}
