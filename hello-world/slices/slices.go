package main

import "fmt"

func main() {
	// Creating slices
	var numbers []int           // nil slice
	scores := []int{95, 87, 92} // slice literal
	grades := make([]int, 3, 5) // make(type, length, capacity)

	// Appending to slices
	numbers = append(numbers, 1, 2, 3)
	scores = append(scores, 88, 91)

	fmt.Printf("Numbers: %v (len=%d, cap=%d)\n", numbers, len(numbers), cap(numbers))
	fmt.Printf("Scores: %v (len=%d, cap=%d)\n", scores, len(scores), cap(scores))
	fmt.Printf("Grades: %v (len=%d, cap=%d)\n", grades, len(grades), cap(grades))

	// Slice operations
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := original[2:7] // Elements 3,4,5,6,7
	slice2 := original[:5]  // Elements 1,2,3,4,5
	slice3 := original[5:]  // Elements 6,7,8,9,10

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Slice1 [2:7]: %v\n", slice1)
	fmt.Printf("Slice2 [:5]: %v\n", slice2)
	fmt.Printf("Slice3 [5:]: %v\n", slice3)

	// Copying slices
	copied := make([]int, len(scores))
	copy(copied, scores)
	fmt.Printf("Copied: %v (len=%d, cap=%d)\n", copied, len(copied), cap(copied))
}
