/**
 * @file: Unit tests for the Deck type
 */

// Package
// *******
package main

// Imports
// *******
import (
	"os"
	"testing"
)

// Test Cases for newDeck()
// ************************
//	- A deck should be created with x number of cards
//	- The 1st card should be an "A of Spade"
//	- The last card should be a "K of Club"
func Test_NewDeck(t *testing.T) {

	// Deck Instance to test on
	var d deck

	// TEST CASE 1: A deck should be created with x number of cards
	// ------------------------------------------------------------
	// Create a new deck
	d = newDeck()

	// Expect: The deck has 52 number of cards
	if len(d) != 52 {
		// If not, something is wrong --> Notify the test-handler t
		t.Errorf("Expected deck length of 52. Got %v", len(d))
	}

	// TEST CASE 2: The 1st card should be an "A of Spade"
	// ---------------------------------------------------
	// Create a new deck
	d = newDeck()

	// Expect: The 1st card is an "A of Spade"
	if d[0] != "A of Spade" {
		// If not, something is wrong --> Notify the test-handler t
		t.Errorf("Expected 1st card to be 'A of Spade'. Got %v", d[0])
	}

	// TEST CASE 3: The last card should be a "K of Club"
	// --------------------------------------------------
	// Create a new deck
	d = newDeck()

	// Expect: The last card is a "K of Club"
	if d[len(d)-1] != "K of Club" {
		// If not, something is wrong --> Notify the test-handler t
		t.Errorf("Expected last card to be 'K of Club'. Got %v", d[0])
	}

}

// Test Cases for saveToFile() and newDeckFromFile()
// -------------------------------------------------
// When testing with files, we have to make sure that we cleanup the files we test with
// Go does not automatically take care of test files
func Test_SaveToFileAndNewDeckFromFile(t *testing.T) {

	// Delete any file _decktesting.tmp from past tests if any
	os.Remove("_decktesting.tmp")

	// Create a new deck
	d := newDeck()

	// Save the deck to file
	d.saveToFile("_decktesting.tmp")

	// Attempt to load from disk
	loaded_deck := newDeckFromFile("_decktesting.tmp")

	// Expect: the Loaded Deck has 52 number of cards
	if len(loaded_deck) != 52 {
		// If not, something is wrong --> Notify the test-handler t
		// Errorf() is a formatted string: We can use % for placeholders
		t.Errorf("Expected loaded deck length of 52. Got %v", len(d))
	}

	// Finally, clean up after ourselves
	os.Remove("_decktesting.tmp")

}

// To run the tests: > go test
// If go.mod is not found in the project dir: > go mod init ./m/v2
// Check go env if needed: > go env
