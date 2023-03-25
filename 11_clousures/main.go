package main

import "fmt"

func main() {
	// EXAMPLE 1:
	increaseTotalAccessCounts := accessCounter()
	fmt.Println(increaseTotalAccessCounts()) // Expected: 1
	fmt.Println(increaseTotalAccessCounts()) // Expected: 2
	fmt.Println(increaseTotalAccessCounts()) // Expected: 3
	fmt.Println(increaseTotalAccessCounts()) // Expected: 4
	fmt.Println(increaseTotalAccessCounts()) // Expected: 5 
	// We can have sure that the value of totalAccess is not loss


	// EXAMPLE 2: 
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	filteredValues := filterNumbers(numbers, func(n int) bool {
		return n % 3 == 0
	})
	fmt.Println(filteredValues)

	// EXAMPLE 3:
	increaseGabrielTotal := increaseByOne(10) // Initial value for Gabriel = 10
	increaseLucasTotal := increaseByOne(0) // Initial value for Lucas = 0

	fmt.Printf("\nGabriel total: %v", increaseGabrielTotal()) // Expected: 11
	fmt.Printf("\nGabriel total: %v", increaseGabrielTotal()) // Expected: 12
	fmt.Printf("\nGabriel total: %v", increaseGabrielTotal()) // Expected: 13

	fmt.Printf("\nLucas total: %v", increaseLucasTotal()) // Expected: 1
	fmt.Printf("\nLucas total: %v", increaseLucasTotal()) // Expected: 2
	fmt.Printf("\nLucas total: %v", increaseLucasTotal()) // Expected: 3

}

// EXAMPLE 1: Clousers allow use to define and access functions within functions 
// and access the variables from internal>external function.
// The fact about it is that the internal vars are preserved
func accessCounter() func() int {
	// This value is "saved"
	totalAccess := 0 

	return func() int {
		totalAccess++
		return totalAccess
	}
}

// EXAMPLE 2: Filter example: Func to filter numbers based into a validation func
// We are passing two params: Array of numbers & filter definition function 
func filterNumbers(numbers[]int, f func(int) bool) []int {
	var filtered []int

	for _, number := range numbers {
		if f(number) {
			filtered = append(filtered, number)
		}
	}

	return filtered
} 

// EXAMPLE3: A fuction in which always increase one depending of a initial number
func increaseByOne(start int) func() int {
	totalValue := start

	return func () int {
		totalValue++
		return totalValue
	}
}