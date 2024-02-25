// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom Interface
// ****************
type ILogWriter struct{}

// ILogWriter implements Writer
func (ILogWriter) Write(bs []byte) (int, error) {
	// Print the byte-slice
	fmt.Println(string(bs))
	// A custom implementation
	fmt.Print("Just wrote this many bytes: ", len(bs))
	// Return
	return len(bs), nil
}

// Functions
// ********
func main() {
	// Create an HTTP request
	resp, err := http.Get("https://example.com")
	// Error Handling
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Using a custom type that implements the Writer interface
	lw := ILogWriter{}
	io.Copy(lw, resp.Body)
}

// FOR WINDOWS:
//	To run:					go run 05-Interfaces\src\main.go
//	To compile:				go build -o 05-Interfaces\bin\main.exe 05-Interfaces\src\main.go
//	To run after compile:	.\05-Interfaces\bin\main.exe
//	Compile + Run:			go build -o 05-Interfaces\bin\main.exe 05-Interfaces\src\main.go && .\05-Interfaces\bin\main.exe

// FOR LINUX:
//	To run:					go run 05-Interfaces/src/main.go
//	To compile:				go build -o 05-Interfaces/bin/main 05-Interfaces/src/main.go
//	To run after compile:	./05-Interfaces/bin/main
//	Compile + Run:			go build -o 05-Interfaces/bin/main 05-Interfaces/src/main.go && ./05-Interfaces/bin/main
