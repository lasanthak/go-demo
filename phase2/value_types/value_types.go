package main

import "fmt"

func modifyValue(x int) {
	x = 100 // This modifies the copy, not the original
}

func modifySlice(s []int) {
	s[0] = 999             // This modifies the original slice's backing array
	s = append(s, 4, 5, 6) // This modifies the copy of the slice header
}

func main() {
	// Value types: int, float, bool, string, arrays, structs
	original := 42
	copy := original
	copy = 50

	fmt.Printf("Original: %d, Copy: %d\n", original, copy) // Original: 42, Copy: 50

	// Function call with value type
	num := 10
	modifyValue(num)
	fmt.Printf("After modifyValue: %d\n", num) // Still 10

	// Reference behavior with slices
	numbers := []int{1, 2, 3}
	fmt.Printf("Before modifySlice: %v\n", numbers)
	modifySlice(numbers)
	fmt.Printf("After modifySlice: %v\n", numbers) // [999, 2, 3] - first element changed
}
