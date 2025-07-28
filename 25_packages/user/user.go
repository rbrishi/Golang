// Package user defines user-related types and functionality
package user

// User is an exported struct type (starts with capital letter)
// It demonstrates:
// 1. Exported struct type with exported fields
// 2. Go's struct definition syntax
// 3. Field naming and visibility
type User struct {
	Name string  // Exported field, accessible from other packages
	Email string // Exported field, accessible from other packages
}
