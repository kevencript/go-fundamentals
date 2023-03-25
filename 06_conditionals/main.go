package main

import (
	"fmt"
)

func main() {
	isRaining := true
	haveUmbrella := true

	// If else
	if isRaining && !haveUmbrella {
		fmt.Println("Is raining outside and you do not have umbrella. You can not go out :(")
	} else if isRaining && haveUmbrella {
		fmt.Println("Is raining but you have your umbrella! You are allowed to go ahead :)")
	} else {
		fmt.Println("Is not raining outside, you can go out :)")
	}

	// Switch case
	age := 22
	switch {
		case age < 18:
			fmt.Println("You definitely can not drive")
		case age > 18 && age < 21:
			fmt.Println("You can not drive yet")
		default:
			fmt.Println("You can drive :)")
	}
}