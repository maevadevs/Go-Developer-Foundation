// PART 2
// ******
// To run this file, rename it to main.go and run as the listed call at the end of the file

// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Functions
// *********

func main() {
	// Grab the filepath from the user
	pathStr := os.Args[1]
	absPath, errorPath := filepath.Abs(pathStr)

	// Check for error
	if errorPath != nil {
		fmt.Println("Error:", errorPath)
		os.Exit(1)
	}

	// Attempt to open the file
	textfilePtr, openfileErr := os.Open(absPath)

	// Check for error
	if openfileErr != nil {
		fmt.Println("Error:", openfileErr)
		os.Exit(1)
	}

	// The File implements the Reader interface
	io.Copy(os.Stdout, textfilePtr)
}

// FOR WINDOWS:
//	To compile:				go build -o 06-Assignments\bin\files.exe 06-Assignments\src\main.go
//	To run after compile:	.\06-Assignments\bin\files.exe 06-Assignments\src\textfile.txt
//	Compile + Run:			go build -o 06-Assignments\bin\files.exe 06-Assignments\src\main.go && .\06-Assignments\bin\files.exe

// FOR LINUX:
//	To compile:				go build -o 06-Assignments/bin/files 06-Assignments/src/main.go
//	To run after compile:	./06-Assignments/bin/files 06-Assignments/src/textfile.txt
//	Compile + Run:			go build -o 06-Assignments/bin/files 06-Assignments/src/main.go && ./06-Assignments/bin/files 06-Assignments/src/textfile.txt
