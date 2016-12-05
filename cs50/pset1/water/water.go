package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

// Function to translate our minutes to "Water Bottles"
func minutesToWaterBottles(min int) int {
	return min * 192 / 16
}

func main() {
	// Thanks Kale Blakenship, and Benjamin Radovsky from the Go Slack
	var minutes int

	// While structure in Go. While minutes is less than 1...
	for minutes < 1 {
		fmt.Printf("minutes: ")

		// TODO: Add tests
		// We grab a new reader and feed it the OS STDIN
		reader := bufio.NewReader(os.Stdin)

		// We read the input from our terminal
		input, _, err := reader.ReadLine()

		if err != nil {
			panic(err)
		}

		// We convert our []byte into a string, and then into an int, base10
		minutes, err = strconv.Atoi(string(input))

		// If we have an error, a negative number, or 0 we keep looping
		if err != nil || minutes < 1 {
			fmt.Println("We need a positive number")
		}
	}

	fmt.Printf("bottles: %d\n", minutesToWaterBottles(minutes))
}
