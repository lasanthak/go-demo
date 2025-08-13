package main

import "fmt"

// Package-level constants
const (
	Pi       = 3.14159
	E        = 2.71828
	MaxRetry = 3
)

// Enumerated constants using iota
const (
	StatusPending  = iota // 0
	StatusActive   = iota // 1
	StatusInactive = iota // 2
	StatusDeleted  = iota // 3
)

const foo string = "foo"
const bar = iota

func main() {
	// Function-level constants
	const greeting = "Hello"
	const port = 8080

	fmt.Printf("Pi: %.2f, E: %.2f\n", Pi, E)
	fmt.Printf("Status: %d\n", StatusInactive)
	fmt.Printf("%s from port %d\n", greeting, port)
	fmt.Println("Foo constant:", foo)
	fmt.Println("Bar constant:", bar)
}
