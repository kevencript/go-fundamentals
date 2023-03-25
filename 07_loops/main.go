package main

import "fmt"

func main() {
	// Long Method
	for i := 1 ; i < 10 ; i++{
		fmt.Print(i)
	}

	// FizzBuzz
	for i := 1 ; i < 100 ; i++{
		switch {
		// Divisible by 3 and 5
		case i % 3 == 0 && i % 5 == 0:
			fmt.Printf("\n%v:FizzBuzz", i)
		
		// Divisible by 3
		case i % 3 == 0 && i % 5 != 0:
			fmt.Printf("\n%v:Fizz", i)

		// Divisible by 5
		case i % 3 != 0 && i % 5 == 0:
			fmt.Printf("\n%v:Buzz", i)
		}
	}
}