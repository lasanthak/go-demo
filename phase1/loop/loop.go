package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Count: %d\n", i)
	}

	// While-style loop
	counter := 0
	for counter < 3 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// Infinite loop with break
	sum := 0
	for {
		sum++
		if sum > 10 {
			break
		}
	}
	fmt.Printf("Sum: %d\n", sum)

	// Range over slice
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Range over map
	ages := map[string]int{"Alice": 30, "Bob": 25}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}
