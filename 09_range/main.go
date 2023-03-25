package main

import (
	"fmt"
)

func main() {
	// Map Loop range
	soccerPlayers := map[string]string {"Messi":"Argentina", "CR7":"Portugal","Neymar":"Brazil"}
	for name, country := range soccerPlayers {
		fmt.Printf("%v play for %v's team!\n", name, country)
	}

	// Slice Loop range
	days := []int {11, 23, 37, 49, 52}
	for _, day := range days { // Ignore the index
		fmt.Printf("\nDay number %v", day)
	}
}