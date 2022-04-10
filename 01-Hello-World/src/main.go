// Package
// *******

package main

// Import
// ******

import "fmt"

// Function
// ********

func main() {

	// Print to the screen
	fmt.Println("Hello World!")

	// Wait before closing the console
	fmt.Println("\nPress enter to exit...")
	fmt.Scanln()

}

// FOR WINDOWS:
// 	To run:					go run 01-Hello-World\src\main.go
// 	To compile:				go build -o 01-Hello-World\bin\program.exe 01-Hello-World\src\main.go
// 	To run after compile:	.\01-Hello-World\bin\program.exe
//	Compile + Run:			go build -o 01-Hello-World\bin\program.exe 01-Hello-World\src\main.go && .\01-Hello-World\bin\program.exe

// FOR LINUX:
// 	To run:					go run 01-Hello-World/src/main.go
// 	To compile:				go build -o 01-Hello-World/bin/program 01-Hello-World/src/main.go
// 	To run after compile:	./01-Hello-World/bin/program
//	Compile + Run:			go build -o 01-Hello-World/bin/program 01-Hello-World/src/main.go && ./01-Hello-World/bin/program
