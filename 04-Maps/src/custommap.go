// Package
// *******
package main

// Import
// ******
import "fmt"

// Type Definitions
// ****************

type custommap map[string]string

// Receiver functions for hashmap
// ******************************

func (cm custommap) print() {
	// Iterating over a map of colors
	for k, v := range cm {
		fmt.Printf("Colors{%s: %s}\n", k, v)
	}
}
