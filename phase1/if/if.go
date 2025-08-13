package main

import "fmt"

func main() {
	age := 25

	// Basic if statement
	if age >= 18 {
		fmt.Println("You are an adult")
	}

	// if-else
	if age < 13 {
		fmt.Println("Child")
	} else if age < 20 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Adult")
	}

	// if with initialization
	if score := 85; score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else {
		fmt.Printf("Grade: %d\n", score)
	}
}
