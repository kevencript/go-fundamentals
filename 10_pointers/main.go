package main

import "fmt"

func main() {
	// Pointer is a var in which have the memory address of the variable in which is pointed to
	
	initialVar := "Initial Text"
	pointer := &initialVar

	// When we access the pointer with *, we access the value of intial variable
	fmt.Printf("Pointer Pure Value (memory address): %v \n Pointer with * (var value): %v", pointer, *pointer)

	// Even if we change the initial var, pointer is able to get the latest assigned value
	initialVar = "Changed text by initial var itself"
	fmt.Printf("\n\nPointer after change: %v", *pointer)

	// Changing the value with the pointer
	*pointer = "Text changed by pointer"
	fmt.Printf("\n\ninitialVar after pointer change: %v\n*pointer after pointer change: %v", initialVar, *pointer)
}