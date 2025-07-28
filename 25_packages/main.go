// Package main is the entry point for the Go application
// In Go, every executable program must have a main package and a main function
package main

// Import section demonstrates different types of imports:
// 1. Standard library imports (fmt)
// 2. External package imports (github.com/fatih/color)
// 3. Local module imports (auth and user packages)
import (
	"fmt" // Standard library for formatting and printing

	"github.com/fatih/color" // Third-party package for colored console output

	// Local module imports using the module path defined in go.mod
	"github.com/rbrishi/Golang/auth" // Local authentication package
	"github.com/rbrishi/Golang/user" // Local user package
)

// Go Modules System:
// 1. Initialize a new module:
//    - Use 'go mod init <module_name>'
//    - Creates go.mod file which tracks dependencies
//    - Module name typically follows GitHub repository path
//
// Example: go mod init github.com/rbrishi/Golang

// main function is the entry point of the program
// It demonstrates:
// 1. Using exported functions from local packages
// 2. Creating and using struct types from local packages
// 3. Using third-party package functionality
func main() {
	// Using exported function from auth package
	// Note: LogInWithCred is capitalized, making it exported (public)
	auth.LogInWithCred("rishabh", "password123")
	fmt.Println(auth.GetSession())

	// Creating an instance of User struct from user package
	// Demonstrates struct initialization with field names
	user := user.User{
		Name: "rishabh",
		Email: "brishabh@gmail.com",
	}
	fmt.Println(user)

	// Using third-party package (color) to print colored text
	// Demonstrates the use of an external package
	color.Red(user.Email)
}


// Go Visibility Rules:
// 1. Exported (Public) names:
//    - Must start with a capital letter
//    - Can be accessed from other packages
//    Example: User, LogInWithCred
//
// 2. Unexported (Private) names:
//    - Start with a lower case letter
//    - Only accessible within the same package
//    Example: forPrivate

// Package Management in Go:
// 1. Finding Packages:
//    - Visit pkg.go.dev (official Go package registry)
//    - Browse and search for packages
//
// 2. Installing Packages:
//    - Use 'go get' command to download and install
//    Example: go get github.com/fatih/color
//
// 3. Dependency Management:
//    - 'go mod tidy' command:
//      * Adds missing dependencies
//      * Removes unused dependencies
//      * Updates go.mod and go.sum files
//      * Ensures dependency graph is complete and accurate


