package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Get initials gets a full name like "Erick Romero"
// And returns initials, "ER", it always capitalizes.
func getInitials(fullname string) string {
	// We separate the names into slices with a space separator
	names := strings.Split(fullname, " ")

	// Result is an empty string at first
	result := ""

	for _, name := range names {
		result += string(name[0])
	}

	return strings.ToUpper(result)
}

func main() {
	var name string

	// While our string is less than 1 char long
	for len(name) < 1 {
		fmt.Println("What's your name?")

		reader := bufio.NewReader(os.Stdin)

		input, _, err := reader.ReadLine()

		if err != nil {
			panic(err)
		}

		name = string(input)
	}

	fmt.Printf("Your initials are %s\n", getInitials(name))
}
