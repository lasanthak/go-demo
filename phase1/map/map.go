package main

import "fmt"

func main() {
	// Creating maps
	var ages map[string]int     // nil map
	ages = make(map[string]int) // initialize

	// Map literal
	cities := map[string]int{
		"New York":    8000000,
		"Los Angeles": 4000000,
		"Chicago":     2700000,
	}

	// Adding and accessing elements
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	fmt.Printf("Ages: %v\n", ages)
	fmt.Printf("Cities: %v\n", cities)

	// Checking if key exists
	age, exists := ages["Alice"]
	if exists {
		fmt.Printf("Alice's age: %d\n", age)
	}

	// Iterating over maps
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// Deleting elements
	delete(ages, "Bob")
	fmt.Printf("After deleting Bob: %v\n", ages)
}
