package main

import (
	"fmt"
)

// Define a structure to store user data
type User struct {
	Name       string
	Email      string
	Environment string
}

func main() {
	// Create a slice to store multiple users
	users := []User{
		{Name: "Alice", Email: "alice@example.com", Environment: "dev"},
		{Name: "Bob", Email: "bob@example.com", Environment: "qc"},
		{Name: "Charlie", Email: "charlie@example.com", Environment: "uat"},
		{Name: "David", Email: "david@example.com", Environment: "prod"},
	}

	// Print the user data
	for _, user := range users {
		fmt.Printf("Name: %s, Email: %s, Environment: %s\n", user.Name, user.Email, user.Environment)
	}
}
