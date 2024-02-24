// Package
// *******

package main

// Import
// ******

import "fmt"

// Structs
// *******

type person struct {
	first_name string
	last_name  string
	contact    contactInfo
}

// Receiver Functions
// ******************

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// Using pointer as receiver allows us to pass-by-reference
//
//	*var - Operator to access the value that exist at the memory address (pointer)
//	*person - A type of Pointer that point to a person type
func (ptr *person) updateName(new_first_name string) {
	(*ptr).first_name = new_first_name
}
