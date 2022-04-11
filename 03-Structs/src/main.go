// Package
// *******

package main

// Import
// ******

import "fmt"

// Structs
// *******

type contactInfo struct {
	email    string
	zip_code int
}

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
//	*var - Operator to access the value that exist at the memory address (pointer)
//	*person - A type of Pointer that point to a person type

func (ptr *person) updateName(new_first_name string) {
	(*ptr).first_name = new_first_name
}

// Helper Functions
// ****************

func updateSlice(s []string) {
	s[0] = "Bye"
}

func updateArray(arr [4]string) {
	arr[0] = "What"
}

// main()
// ******

func main() {

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
// 	To run:					go run src\main.go
// 	To compile:				go build -o bin\program.exe src\main.go
// 	To run after compile:	.\bin\program.exe

// FOR LINUX: FROM PROJECT FOLDER
// 	To run:					go run src/main.go
// 	To compile:				go build -o bin/program src/main.go
// 	To run after compile:	./bin/program
//	Compile + Run:			go build -o bin/program src/main.go && ./bin/program
