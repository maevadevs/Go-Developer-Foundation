// Package
// *******

package main

// Types
// *****

type (
	englishBot struct{}
	spanishBot struct{}
)

// Receiver Functions
// ******************

func (englishBot) getGreeting() string {
	// Imagine some logic that is custom to englishBot
	// Custom implementation for englishBot
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	// Imagine some logic that is custom to spanishBot
	// Custom implementation for spanishBot
	return "Hola!"
}
