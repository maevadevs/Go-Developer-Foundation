// Package
// *******
package main

// Import
// ******
import "fmt"

// Type Definitions
// ****************
type hashmap map[string]string

// Receiver functions for hashmap
// ******************************
func (c hashmap) print() {
	// Iterating over a map of colors
	for color, hex := range c {
		fmt.Printf("Colors{%s: %s}\n", color, hex)
	}
}
