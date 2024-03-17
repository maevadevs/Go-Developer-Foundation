/**
 * @file: Unit tests for the Deck type
 * To run the tests, make sure to be in the same location as this file
 * In VS Code, we have the options for "Run Package Tests" and "Run File Tests"
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
//   - A deck should be created with x number of cards
//   - The 1st card should be an "A of Spade"
//   - The last card should be a "K of Club"
func Test_newDeck(t *testing.T) {
	// Deck Instance to test on
	var d deck

	// TEST CASE 1: A deck should be created with x number of cards
	// ------------------------------------------------------------
	// Create a new deck
	d = newDeck()

	// Set expectations:
	// The deck should have 52 number of cards
	expectedDeckLen := 52
	actualDeckLen := len(d)

	// Test expectations
	if expectedDeckLen != actualDeckLen {
		// If not, something is wrong --> Notify the test-handler t
		t.Errorf("Test Case 1: Expected deck length of %v. Got %v", expectedDeckLen, actualDeckLen)
	}

	// TEST CASE 2: The 1st card should be an "A of Spade"
	// ---------------------------------------------------
	// Create a new deck
	d = newDeck()

	// Set expectations:
	// The 1st card is "A of Spade"
	expected1stCard := "A of Spade"
	actual1stCard := d[0]

	// Test expectations
	if expected1stCard != actual1stCard {
		// If not, something is wrong --> Notify the test-handler t
		t.Errorf("Expected 1st card to be %v. Got %v", expected1stCard, actual1stCard)
	}

	// TEST CASE 3: The last card should be a "K of Club"
	// --------------------------------------------------
	// Create a new deck
	d = newDeck()

	// Set expectations:
	// The last card is "K of Club"
	expectedLastCard := "K of Club"
	actualLastCard := d[len(d)-1]

	// Test expectations
	if expectedLastCard != actualLastCard {
		// If not, something is wrong --> Notify the test-handler t
		t.Errorf("Expected last card to be %v. Got %v", expectedLastCard, actualLastCard)
	}
}

// Test Cases for saveToFile() and newDeckFromFile()
// -------------------------------------------------
// When testing with files, we have to make sure that we cleanup the files we test with
// Go does not automatically take care of test files
func Test_saveToFileAndNewDeckFromFile(t *testing.T) {
	// Delete any file _decktesting.tmp from past tests if any
	os.Remove("_decktesting.tmp")

	// Create a new deck
	d := newDeck()

	// Save the deck to file
	d.saveToFile("_decktesting.tmp")

	// Attempt to load from disk
	loadedDeck := newDeckFromFile("_decktesting.tmp")

	// Set expectations:
	// The length of the loaded deck from file should be the same as the original deck
	expectedLoadedDeckLen := len(d)
	actualLoadedDeckLen := len(loadedDeck)

	// Test expectations
	if expectedLoadedDeckLen != actualLoadedDeckLen {
		// If not, something is wrong --> Notify the test-handler t
		// Errorf() is a formatted string: We can use % for placeholders
		t.Errorf("Expected loaded deck length of %v. Got %v", expectedLoadedDeckLen, actualLoadedDeckLen)
	}

	// Finally, clean up any temp test files
	os.Remove("_decktesting.tmp")
}

// To run the tests: > go test ./src
// If go.mod is not found in the project dir: > go mod init ./m/v2
// Check go env if needed: > go env
