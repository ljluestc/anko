package main

import (
	"fmt"
	"log"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/vm"
)

// User is a test struct with a pointer receiver method.
type User struct {
	Name string
}

func (u *User) SayHello() {
	fmt.Printf("Hello %s!\n", u.Name)
}

func main() {
	// Create a new Anko environment.
	e := env.NewEnv()

	// Define a User instance (struct value, not pointer).
	user := User{Name: "Michel"}

	// Verify the method works in Go.
	user.SayHello() // Expected: Hello Michel!

	// Define the user in Anko's environment.
	err := e.Define("user", user)
	if err != nil {
		log.Fatalf("Define error: %v", err)
	}

	// Execute the method call in Anko's VM.
	_, err = vm.Execute(e, nil, "user.SayHello()") // Remove unused result variable
	if err != nil {
		log.Fatalf("Execution error: %v", err)
	}

	// Expected output:
	// Hello Michel!
	// Hello Michel!
}
