package main

import (
	"fmt"
)

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr (only postive numbers)
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Testing vars and const (same as ES6 JS concepts)
	var name = "Gabriel"
	var age byte = 37
	const isCool = true

	// ShortHand
	shortHandString := "testing var"
	size := 1.3

	// Multiple ShortHand assignments
	city, year := "brazil", 1999


	// Reference: https://pkg.go.dev/fmt
	fmt.Println(name, age, shortHandString, isCool, city, year, size)
	fmt.Printf("How Go see it: %#v\n Data type: %T\n", name, age)
}

