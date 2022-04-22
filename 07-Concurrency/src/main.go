// Package
// *******
package main

// Imports
// *******
import (
	"fmt"
	"net/http"
	"time"
)

// Functions
// ********
func main() {

	// A list of urls
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	// Channel for communicating with go routines
	ch := make(chan string)

	// Loop through the links
	for _, url := range urls {
		// Create a new Go Routine for each call
		go checkLink(url, ch)
	}

	for l := range ch {
		// Receive message from the channel
		go func(link string) {
			// time.sleep pauses the current Go Routine
			time.Sleep(time.Second * 2) // Pause of 2 seconds
			// Then span a new go routine to recheck the link again
			checkLink(link, ch)
		}(l)
	}

}

// Helper Functions
// ****************
func checkLink(link string, ch chan string) {

	// Test the link with a Get call
	_, err := http.Get(link)

	// If error, then we have an issue
	if err != nil {
		fmt.Println(link, "might be down")
		// Send message via the channel
		ch <- link
		return
	}

	// Else, we are good
	fmt.Println(link, "is up")
	// Send message via the channel
	ch <- link

}

// FOR WINDOWS:
//	To run:					go run 07-Concurrency\src\main.go
//	To compile:				go build -o 07-Concurrency\bin\concurrency.exe 07-Concurrency\src\main.go
//	To run after compile:	.\07-Concurrency\bin\concurrency.exe
//	Compile + Run:			go build -o 07-Concurrency\bin\concurrency.exe 07-Concurrency\src\main.go && .\07-Concurrency\bin\concurrency.exe

// FOR LINUX:
//	To run:					go run 07-Concurrency/src/main.go
//	To compile:				go build -o 07-Concurrency/bin/concurrency 07-Concurrency/src/main.go
//	To run after compile:	./07-Concurrency/bin/concurrency
//	Compile + Run:			go build -o 07-Concurrency/bin/concurrency 07-Concurrency/src/main.go && ./07-Concurrency/bin/concurrency
