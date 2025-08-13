package main

import "fmt"

func main() {
	// Various ways to declare variables
	var name string = "John" // Explicit type
	var age = 30             // Type inference
	var isActive bool        // Zero value (false)

	// Short variable declaration (inside functions only)
	city := "Austin"
	temperature := 98.6

	// Multiple variable declaration
	var (
		firstName string = "Jane"
		lastName  string = "Doe"
		salary    int    = 75000
	)

	fmt.Printf("Name: %s, Age: %d, Active: %t\n", name, age, isActive)
	fmt.Printf("City: %s, Temp: %.1f\n", city, temperature)
	fmt.Printf("Employee: %s %s, Salary: $%d\n", firstName, lastName, salary)
}
