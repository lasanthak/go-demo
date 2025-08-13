package main

import "fmt"

func main() {
	// Basic pointer operations
	x := 42

	// Get pointer to x
	ptr := &x

	// Dereference pointer
	value := *ptr

	fmt.Printf("x = %d\n", x)         // x = 42
	fmt.Printf("x @ %p\n", &x)        // x = 0x... (memory address)
	fmt.Printf("ptr = %p\n", ptr)     // ptr = 0x... (memory address)
	fmt.Printf("*ptr = %d\n", *ptr)   // *ptr = 42
	fmt.Printf("value = %d\n", value) // value = 42

	// Modify through pointer
	*ptr = 100
	fmt.Printf("After *ptr = 100, x = %d, value = %d\n", x, value) // x = 100, value = 42

	// Pointer to pointer
	ptrToPtr := &ptr
	fmt.Printf("ptrToPtr = %p\n", ptrToPtr)
	fmt.Printf("*ptrToPtr = %p\n", *ptrToPtr)
	fmt.Printf("**ptrToPtr = %d\n", **ptrToPtr)
}
