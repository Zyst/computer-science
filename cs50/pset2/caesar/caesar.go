package main

import (
	"fmt"
)

// Does the caesar cipher letter shift
func caesarShift(letter byte, offset int) byte {
	// If we have an uppercase letter or lowercase letter
	if letter >= 'A' && letter <= 'Z' {
		return (((letter - 'A') + byte(offset)) % 26) + 'A'
	} else if letter >= 'a' && letter <= 'z' {
		return (((letter - 'a') + byte(offset)) % 26) + 'a'
	}
	// If we got to this point we return the character unchanged
	return letter
}

// Iterates through our input string and runs it through the
// caesar shift function above
func caesarCipher(cipher string, offset int) string {
	result := ""

	for _, s := range cipher {
		result += string(caesarShift(byte(s), offset))
	}

	return result
}

func main() {
	fmt.Printf("%c\n", caesarShift('y', 2))
}
