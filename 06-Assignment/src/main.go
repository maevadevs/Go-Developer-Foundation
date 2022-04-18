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
// ********
func main() {
	// Grab the filepath from the user
	path_str := os.Args[1]
	abs_path, _ := filepath.Abs(path_str)
	// Attempt to open the file
	textfile_ptr, openfile_err := os.Open(abs_path)
	// Check for error
	if openfile_err != nil {
		fmt.Println("Error:", openfile_err)
		os.Exit(1)
	}
	// The File implements the Reader interface
	io.Copy(os.Stdout, textfile_ptr)
}

// FOR WINDOWS:
//	To compile:				go build -o 06-Assignment\bin\files.exe 06-Assignment\src\main.go
//	To run after compile:	.\06-Assignment\bin\files.exe 06-Assignment\src\textfile.txt
//	Compile + Run:			go build -o 06-Assignment\bin\files.exe 06-Assignment\src\main.go && .\06-Assignment\bin\files.exe

// FOR LINUX:
//	To compile:				go build -o 06-Assignment/bin/files 06-Assignment/src/main.go
//	To run after compile:	./06-Assignment/bin/files textfile.txt
//	Compile + Run:			go build -o 06-Assignment/bin/files 06-Assignment/src/main.go && ./06-Assignment/bin/files textfile.txt
