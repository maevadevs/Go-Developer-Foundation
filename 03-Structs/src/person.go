// Package
// *******

package main

// Import
// ******

import "fmt"

// Structs
// *******

type person struct {
	firstName string
	lastName  string
	// Embedded struct
	contact contactInfo
}

// Receiver Functions
// ******************

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// Using pointer as receiver allows us to pass-by-reference
//
//	*var - Operator to access the value that exist at the memory address (pointer)
//	       But Go can also automatically de-reference
//	*person - A type of Pointer that point to a person type
func (ptrPers *person) updateFirstName(newFirstName string) {
	ptrPers.firstName = newFirstName
}
