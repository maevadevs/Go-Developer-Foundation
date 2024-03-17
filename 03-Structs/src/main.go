// Package
// *******

package main

// Import
// ******

import "fmt"

// main()
// ******
func main() {
	// Declare variable:
	// Fields default to null-value
	var julie person

	// To print a struct with key:value format, use %+v
	fmt.Printf("julie before assignment: %+v\n", julie)

	// Assigning / Re-assigning values
	julie.firstName = "Julie"
	julie.lastName = "Arkorius"
	julie.contact.email = "j.arkorius@somemail.com"
	julie.contact.zipCode = 12345

	// Check again
	fmt.Printf("julie after assignment: %+v\n", julie)

	fmt.Println()

	// Creating a new person
	jim := person{
		firstName: "Jim",
		lastName:  "Patterson",
		// Embedded struct
		contact: contactInfo{
			email:   "jim@patterson.com",
			zipCode: 98765,
		},
	}

	// Calling a receiver function
	jim.print()
	(&jim).updateFirstName("Jimmy")
	jim.print()

	// We can call with either 'jim' or '&jim'
	// Go is able to automatically reference and de-reference a pointer
	// As long as the function's parameter is defined correctly as pointer
	// With Go, it is possible to substitute a pointer with its root variable
	// So the following still work, even if the receiver requires a pointer type
	jim.updateFirstName("Big ol'Jim")
	jim.print()

	fmt.Println("--------------------------------------------")

	mySlice := []string{"Hi", "there", "how", "are", "you"}
	myArray := [4]string{"This", "is", "an", "Array"}

	updateSlice(mySlice)
	updateArray(myArray)

	fmt.Println(mySlice)
	fmt.Println(myArray)
}

// FOR WINDOWS: FROM PROJECT FOLDER
//  To compile:             go build -o 03-Structs\bin\Program.exe .\03-Structs\src\
//  To run after compile:   .\03-Structs\bin\Program.exe

// FOR LINUX: FROM PROJECT FOLDER
//  To compile:             go build -o 03-Structs/bin/Program ./03-Structs/src/
//  To run after compile:   ./03-Structs/bin/Program
//  Compile + Run:          go build -o 03-Structs/bin/Program ./03-Structs/src/ && ./03-Structs/bin/Program
