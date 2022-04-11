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

// FOR WINDOWS: FROM PROJECT FOLDER
// 	To run:					go run src\main.go
// 	To compile:				go build -o bin\program.exe src\main.go
// 	To run after compile:	.\bin\program.exe

// FOR LINUX: FROM PROJECT FOLDER
// 	To run:					go run src/main.go
// 	To compile:				go build -o bin/program src/main.go
// 	To run after compile:	./bin/program
//	Compile + Run:			go build -o bin/program src/main.go && ./bin/program
