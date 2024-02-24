// Package
// *******

package main

// Import
// ******

import "fmt"

// main()
// ******
func main() {
	// Creating a new person
	jim := person{
		first_name: "Jim",
		last_name:  "Patterson",
		contact: contactInfo{
			email:    "jim@patterson.com",
			zip_code: 98765,
		},
	}

	// Calling a receiver function
	jim.print()
	(&jim).updateName("Jimmy")
	jim.print()

	// With Go, it is possible to substitute a pointer with its root variable
	// So the following still work, even if the receiver requires a pointer type
	jim.updateName("Big ol'Jim")
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
// 	To compile:				go build -o bin\program.exe .\src\
// 	To run after compile:	.\bin\program.exe

// FOR LINUX: FROM PROJECT FOLDER
// 	To compile:				go build -o bin/program ./src/
// 	To run after compile:	./bin/program
//	Compile + Run:			go build -o bin/program ./src/ && ./bin/program
