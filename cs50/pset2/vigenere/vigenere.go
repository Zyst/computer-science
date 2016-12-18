package main

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

		// If our offset is lowercase we turn it into uppercase
		if isLower(offset) {
			offset -= 32
		}

		return ((letter - 'A' + offset - 'A') % 26) + 'A'
	} else if isLower(letter) {

		// If our offset is uppercase we turn it into lowercase
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
