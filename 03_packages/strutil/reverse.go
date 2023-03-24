package strutil

// NICE TO KNOW:
// 1. Functions starting with uppercase letter are visible OUTSIDE the package
// 2. Funcs starting with lowercase level are visible only WITHIN the package

func ReverseString(s string) string {
	// The slice [71 97 98 114 105 101 108] represents a sequence of integers
	// that correspond to the Unicode values (also known as code points) of
	// the characters in a string. In Go, these Unicode values are called "runes".

	// In this case, the rune sequence represents the string "Gabriel".
	// Here's the mapping of Unicode values to characters:
	// - 71: 'G'
	// - 97: 'a'
	// - 98: 'b'
	// - 114: 'r'
	// - 105: 'i'
	// - 101: 'e'
	// - 108: 'l'

	// When working with strings in Go, it can sometimes be helpful to convert
	// them into a rune slice, especially when dealing with Unicode characters.
	// This allows you to correctly process and manipulate individual characters,
	// even if they may have varying byte sizes.

	// For example, you can convert a string into a rune slice like this:
	// s := "Gabriel"
	// runes := []rune(s)

	// Now you can manipulate the `runes` slice as you wish,
	// and then convert it back to a string:
	// s = string(runes)
	runes := []rune(s)

	// Initialize two pointers: one at the start and one at the end of the rune slice
	start, end := 0, len(runes)-1

	// Loop through the rune slice, swapping the runes at the start and end pointers
	// Move the pointers towards each other until they meet in the middle
	for start < end {
		// Swap the runes at the start and end pointers
		runes[start], runes[end] = runes[end], runes[start]

		// Move the start pointer one step forward and the end pointer one step backward
		start++
		end--
	}

	// Convert the rune slice back to a string and return it
	return string(runes)
}

