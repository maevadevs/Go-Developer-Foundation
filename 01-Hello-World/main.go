// Package
// *******
// Package == Project == Workspace
// - A Package can have many files associated with it
// - Each file ending with .go
// - The very first line of each file in the same package must start with the package declaration

// There are 2 types of packages:
//	1. Executable
//		- Always "main" package
//		- Must always have a function called "main()" as well
//		- Generates a .exe/.bin file that we can run when compiled
//		- Code used when we want to do something (Executed codes)
//	2. Reusable
//		- Any other package name other than "main"
//		- Code used as "helpers" for reusable logic, libraries, dependencies
// - Files in the same package do not have to be imported into each other before they can be used
// - If other packages are declared with `package main`, the functions they contains can be used here directly
// - But one of the file must contain the the "main()" function as the primary entry

package main

// Import
// ******
// Allows to import codes from other packages (Reusable)
// "fmt"
//	- A standard package within Go (pkg.go.dev/std)
//	- Short for "format"
//	- Mostly used for debugging and development
// Unless we import a package, we have no access to any functionalities in another package
// We are not limited to import packages from the standard library
//	- We can import packages written by other engineers as well

import "fmt"

// Function
// ********
// A "main" package must have a main() function as the entry-point
// This is a Go func, similar functionality to other languages

func main() {
	fmt.Println("Hello World!")
}

// How is a Go file organized?
// ***************************
// In practice, always the same typical pattern
//
// 1. Package
// 2. Imports
// 3. Functions

/* How do we run code in our project?
 * **********************************
 * > go run <filename>
 * > go run "<filepath>"
 */

/* Available Go Commands
 * *********************
 * > go build 	- Compiles a bunch of go source code files into executable binaries
 * > go run 	- Compiles and execute a bunch of go source code files (build + run) but does not produce an actual executable
 * > go fmt		- Formats all the code in each file in the current directory
 * > go install	- Compiles and "install" a package
 * > go get		- Download the raw source code of someone else's package
 * > go test	- Runs any tests associated with the current projects
 */
