package main

import "fmt"

// incrementValue is a function that takes a pointer to an integer,
// increments the value it points to, and prints the address and new value.
func incrementValue(valPtr *int) {
	fmt.Println("\nInside incrementValue function (pointer version):")
	fmt.Printf("Address of variable inside function: %p\n", valPtr)
	fmt.Printf("Address where the value is stored (dereferenced pointer): %p\n", &*valPtr)
	fmt.Printf("Value before increment: %d\n", *valPtr)

	*valPtr++

	fmt.Printf("Value after increment: %d\n", *valPtr)
	fmt.Printf("Address of variable inside function after increment (still same pointer address): %p\n", valPtr)
	fmt.Printf("Address where the value is stored after increment (still same memory location): %p\n", &*valPtr)
}

// incrementValueNoPtr is a function that takes an integer by value,
// increments it, and prints the address and new value.
// IMPORTANT: This function operates on a COPY of the original 'count' variable.
func incrementValueNoPtr(val int) {
	fmt.Println("\nInside incrementValueNoPtr function (no pointer version):")
	fmt.Printf("Address of variable inside function: %p\n", &val) // Address of the copy 'val'
	fmt.Printf("Value before increment: %d\n", val)

	val++ // Increment the COPY of the value

	fmt.Printf("Value after increment: %d\n", val)
	fmt.Printf("Address of variable inside function after increment (still same address of copy): %p\n", &val)
}

func main() {
	// 1. Declare a variable with a hardcoded value
	count := 10

	// 2. Show initial address and value
	fmt.Println("Initial state:")
	fmt.Printf("Variable name: count\n")
	fmt.Printf("Address of count in memory: %p\n", &count)
	fmt.Printf("Value of count: %d\n", count)

	// 3. Call the increment function (pointer version), passing the address of 'count'
	incrementValue(&count)

	// 4. Show address and value after incrementing (pointer version)
	fmt.Println("\nAfter incrementValue function (pointer version):")
	fmt.Printf("Address of count in memory (after incrementValue): %p\n", &count)
	fmt.Printf("Value of count (after incrementValue): %d\n", count)

	// 5. Call the increment function (no pointer version), passing the value of 'count'
	incrementValueNoPtr(count) // Passing the VALUE of 'count'

	// 6. Show address and value after incrementing (no pointer version)
	fmt.Println("\nAfter incrementValueNoPtr function (no pointer version):")
	fmt.Printf("Address of count in memory (after incrementValueNoPtr): %p\n", &count) // Address should remain the same as before incrementValueNoPtr
	fmt.Printf("Value of count (after incrementValueNoPtr): %d\n", count)              // Value should NOT be changed by incrementValueNoPtr
}
