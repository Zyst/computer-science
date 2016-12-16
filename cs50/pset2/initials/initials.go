package main

import (
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

}
