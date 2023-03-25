package main

import "fmt"

func main() {

	// Define map
	// OBS: Needs make() because we are creating emails and not assigning anything 
	emails := make(map[string]interface{})

	// Define emails 
	emails["Bob"] = "bob@mail.com"
	emails["Junior"] = 124 // test
	emails["Pedro"] = "Pedro@mail.com"

	fmt.Printf("Complete map: %v", emails)
	fmt.Printf("\nTotal map length: %v", len(emails))
	

	// Validating if a key exists
	mapValue, exists := emails["Junior"]  
	  // mapValue = string (or the type of map value)
	  // exists   = bool

	fmt.Printf("\nDo 'Junior' exists? R: Exists: %v, Value: %v ", exists, mapValue)

	// Deleting values from the map
	delete(emails, "Junior")
	fmt.Printf("\nUpdated map: %v", emails)
	newMapValue, exists := emails["Junior"]  
	fmt.Printf("\nDo 'Junior' exists? R: Exists: %v, Value: %v ", exists, newMapValue)
	fmt.Printf("\nTotal map length: %v", len(emails))


	// Crating and assigning map values instantly
	// OBS: Do not need make()
	soccerPlayers := map[string]string {"Messi":"Argentina", "CR7":"Portugal","Neymar":"Brazil",}
	fmt.Printf("\nSoccer players: %v", soccerPlayers)
}