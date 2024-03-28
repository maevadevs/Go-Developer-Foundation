// PART 4
// ******
// To run this file, rename it to main.go and run as the listed call at the end of the file

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
	// A slice of urls
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://go.dev",
		"https://amazon.com",
	}

	// Channel for communicating with go routines
	ch := make(chan string)

	// Loop through the urls
	for _, url := range urls {
		// Create a new Go Routine for each call
		// Pass the channel to the new routine
		go checkUrl(url, ch)
	}

	// Receive the messages from the channel
	// NOTE: This code is blocking
	// The main routine will wait here until something happen
	// Once something happen, it continues execution
	// fmt.Println(<-ch)

	// So we would need to check the channel multiple times for each urls
	// for i := 0; i < len(urls); i++ {
	// 	// Receive message from the channel
	// 	// This is a blocking call
	// 	fmt.Println(<-ch)
	// }

	// Keep on checking the url in an infinite loop
	// for {
	// 	// Receive message from the channel
	//     // Span a new go routine to recheck the url again
	//     // This is a blocking call
	// 	go checkUrl(<-ch, ch)
	// }

	// Alternative Syntax: `range` can also be used with channels
	// Keep on checking the url in an infinite loop
	// for l := range ch {
	//     // Receive message from the channel
	//     // Span a new go routine to recheck the url again
	//     // This is a blocking call
	//     go checkUrl(l, ch)
	// }

	// We should add a slight pause between each new call
	// Keep on checking the url in an infinite loop
	for l := range ch {
		// Receive message from the channel
		// This is a blocking call
		// Use a function literal so to not block the main routine
		go func(lnk string) {
			// NOTE: time.Sleep pauses the current Go Routine, which is the function literal
			// Pause of 2 seconds
			time.Sleep(time.Second * 2)

			// Recheck the url again
			checkUrl(lnk, ch)
		}(l)
	}
}

// Helper Functions
// ****************

// Check if a URL is reachable or not.
func checkUrl(url string, ch chan string) {
	// Test the url with a Get call
	_, err := http.Get(url)
	// Error Handling
	if err != nil {
		fmt.Println(url, "might be down")
		// Send the url via the channel
		ch <- url
		return
	}

	// Else, we are good
	fmt.Println(url, "is up")

	// Send the url via the channel
	ch <- url
}

// FOR WINDOWS:
//  To run:                 go run 07-Concurrency\src\main.go
//  To compile:             go build -o 07-Concurrency\bin\concurrency.exe 07-Concurrency\src\main.go
//  To run after compile:   .\07-Concurrency\bin\concurrency.exe
//  Compile + Run:          go build -o 07-Concurrency\bin\concurrency.exe 07-Concurrency\src\main.go && .\07-Concurrency\bin\concurrency.exe

// FOR LINUX:
//  To run:                 go run 07-Concurrency/src/main.go
//  To compile:             go build -o 07-Concurrency/bin/concurrency 07-Concurrency/src/main.go
//  To run after compile:   ./07-Concurrency/bin/concurrency
//  Compile + Run:          go build -o 07-Concurrency/bin/concurrency 07-Concurrency/src/main.go && ./07-Concurrency/bin/concurrency
