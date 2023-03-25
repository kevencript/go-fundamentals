package main

import "fmt"

// Here we can create structs
type Car struct {
	name, energyType string
	year int
	isTurnedOn bool
}

// Defining methods for structs

// Methods in which DO NOT change anything into the struct:
// obs: is not necessary to specify the pointer * 
func (c Car) returnCarSummary() map[string]interface{} {
	return map[string]interface{} {
		"name": c.name,
		"year": c.year,
		"isTurnedOn": c.isTurnedOn,
	} 
}

// Methods in which change anything into the struct:
// obs: is necessary to specify the pointer * 
func (c *Car) turnOn() {
	fmt.Println("Vrum Vrum... Car is turned on!")
	if !c.isTurnedOn {
		c.isTurnedOn = true
	}
}


func main() {
	// Creating example car and accessing those methods
	newCar := Car{name: "BMW", energyType: "Battery", year: 1990, isTurnedOn: false}
	fmt.Println(newCar.returnCarSummary())
	newCar.turnOn()
	fmt.Println(newCar.returnCarSummary())
}