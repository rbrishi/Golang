// Package auth (session.go)
// This file is part of the auth package and shares the same package declaration
package auth

// GetSession is an exported function that returns the current session status
// It demonstrates:
// 1. Package member visibility (exported function)
// 2. Multiple files in the same package
// 3. Package-level organization of related functionality
func GetSession() string {
	return "session loged in"
}