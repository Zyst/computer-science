package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var encrypt string
	offset, err := strconv.ParseInt(os.Args[1], 10, 64)

	if err != nil {
		panic("You must use an int as a command line argument")
	}

	// While our string is less than 1 char long
	for len(encrypt) < 1 {
		fmt.Println("Encrypt the following:")

		reader := bufio.NewReader(os.Stdin)

		input, _, err := reader.ReadLine()

		if err != nil {
			panic(err)
		}

		encrypt = string(input)
	}

	fmt.Printf("Your encrypted string is: %s\n", caesarCipher(encrypt, int(offset)))
}
