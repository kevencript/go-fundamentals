package main

import (
	"fmt"
	"math"

	"github.com/kevencript/go-fundamentals/03_packages/strutil"
)

func main() {
	fmt.Printf("Math package: %v\n", math.Ceil(1.9))
	fmt.Printf("Own package (strutil): %v ", strutil.ReverseString("Gabriel"))
}	