package main

import "fmt"

func main() {
	// Arrays 1: Assign values separately
	var fruitArr [2]string 

	fruitArr[0] = "Orange"
	fruitArr[1] = "Apple"

	fmt.Printf("Manual Array: %s\n", fruitArr)

	// Arrays 2: Assign value automatically
	carsArr := [2]string{"BMW", "Mercedes"}

	fmt.Printf("Auto Array  : %v\n", carsArr)

	// Slices (Same, but without specify the number within [] brackets)
	citiesArr := []string {"Brasilia", "Sergipe", "Brasília"}

	fmt.Printf("Slices: %v", citiesArr)

	// Mixed Slice (multiple datatypes)
	randomArr := []interface{} {"Array", 1.2, -1, 7, [2]string {"test1", "test2"}}

	fmt.Println("\nMixed slices: ")
	for _, v := range randomArr {
		fmt.Printf("\n- Name: %v\n - Type: %T\n", v, v)
	}

	// len() - get length
	fmt.Printf("\nCars Length: %v", len(carsArr))
}