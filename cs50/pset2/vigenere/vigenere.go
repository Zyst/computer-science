package main

import (
	"fmt"
)

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// Does the Vigenere cipher letter shift
func vigenereShift(letter byte, offset byte) byte {

	// If we have an uppercase letter or lowercase letter
	if isUpper(letter) {

    // If our offset is lowercase
    if isLower(offset) {
      offset -= 32
		}

		return ((letter - 'A' + offset - 'A') % 26) + 'A'
	} else if isLower(letter) {

    // If our offset is lowercase
    if isUpper(offset) {
      offset += 32
		}

		return ((letter - 'a' + offset - 'a') % 26) + 'a'
	}
	// If we got to this point we return the character unchanged
	return letter
}

func vigenereCipher(encrypt string, key string) string {
	result := ""

	// Key index for vigenere
	for i, s := range encrypt {
		fmt.Printf("%c --- %c and key[%d]\n", vigenereShift(byte(s), key[i%5]), s, i%5)
	}

	return result
}

func main() {

}
