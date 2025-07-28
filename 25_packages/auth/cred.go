// Package auth provides authentication-related functionality
// It demonstrates Go's package system and visibility rules
package auth

import "fmt"

// LogInWithCred is an exported function (starts with capital letter)
// It handles user authentication with username and password
// Parameters:
//   - username: user's login name
//   - password: user's password
func LogInWithCred(username, password string)  {
	forPrivate() // Calling an unexported function within the same package
	fmt.Println("Logging in with credentials:", username, password)
}

// forPrivate is an unexported function (starts with lowercase letter)
// It can only be used within the auth package
// This demonstrates Go's visibility rules for package members
func forPrivate() string {
	return ""
}