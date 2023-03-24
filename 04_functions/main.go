package main

import (
	"fmt"
)

func greeting(name, country string) string {
	return "Hello " + name + " from " + country
}

func main() {
	fmt.Print(greeting("Gabriel", "New Island"))
}	