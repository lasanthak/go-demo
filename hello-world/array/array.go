package main

import "fmt"

func main() {
	// Array declaration and initialization
	var numbers [5]int                 // Zero-valued array
	scores := [3]int{95, 87, 92}       // Array literal
	grades := [...]int{85, 90, 78, 88} // Compiler counts elements

	// Accessing and modifying arrays
	numbers[0] = 10
	numbers[1] = 20

	fmt.Printf("Numbers: %v\n", numbers)
	fmt.Printf("Scores: %v\n", scores)
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Array length: %d\n", len(grades))

	// Iterating over arrays
	for i, value := range scores {
		fmt.Printf("Index %d: %d\n", i, value)
	}
}
