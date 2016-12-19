package main

// Checks if a letter is uppercase
func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

// Checks if a letter is lowercase
func isLower(b byte) bool {
	return b >= 'a' && b <= 'z'
}

// Does the Vigenere cipher letter shift
func vigenereShift(letter byte, offset byte) byte {

	// If we have an uppercase letter or lowercase letter
	if isUpper(letter) {

		// If our offset is lowercase we turn it into uppercase
		// because Vigenere keys are case insensitive
		if isLower(offset) {
			offset -= 32
		}

		// We return the letter turned into an alphabetical index (0-25 A-Z) plus
		// the offset turned into an alphabetical index, then we do a modulo
		// 26 operation, and finally turn the letter into the ASCII range again
		return ((letter - 'A' + offset - 'A') % 26) + 'A'
	} else if isLower(letter) {

		// If our offset is uppercase we turn it into lowercase
		// because Vigenere keys are case insensitive
		if isUpper(offset) {
			offset += 32
		}

		// We return the letter turned into an alphabetical index (0-25 a-z) plus
		// the offset turned into an alphabetical index, then we do a modulo
		// 26 operation, and finally turn the letter into the ASCII range again
		return ((letter - 'a' + offset - 'a') % 26) + 'a'
	}
	// If we got to this point we return the character unchanged
	return letter
}

// We iterate through our string, running our vignere shift algorithm on each
// individual letter, and appending it to a result string as appropriate
func vigenereCipher(encrypt string, key string) string {
	result := ""

	// The sub-index we use for our vigenere key
	subIndex := 0

	// Key index for vigenere
	for _, s := range encrypt {
		r := vigenereShift(byte(s), key[subIndex%5])

		// If our value is a letter we increase our subindex
		if isUpper(byte(s)) || isLower(byte(s)) {
			subIndex++
		}

		// We append to result
		result += string(r)
	}

	return result
}
