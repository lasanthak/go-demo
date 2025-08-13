package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Custom error type
type ValidationError struct {
	Field string
	Value string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s' with value '%s'", e.Field, e.Value)
}

// Function that returns custom error
func validateAge(ageStr string) (int, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, ValidationError{Field: "age", Value: ageStr}
	}

	if age < 0 || age > 150 {
		return 0, errors.New("age must be between 0 and 150")
	}

	return age, nil
}

func main() {
	// Error handling examples
	testValues := []string{"25", "abc", "-5", "200"}

	for _, value := range testValues {
		age, err := validateAge(value)
		if err != nil {
			fmt.Printf("Error with '%s': %v\n", value, err)

			// Type assertion for custom errors
			if ve, ok := err.(ValidationError); ok {
				fmt.Printf("  Custom error - Field: %s, Value: %s\n", ve.Field, ve.Value)
			}
		} else {
			fmt.Printf("Valid age: %d\n", age)
		}
	}
}
