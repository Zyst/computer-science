package main

import (
	"fmt"
)

// Function to translate our minutes to "Water Bottles"
func minutesToWaterBottles(min int) int {
	return min * 192 / 16
}

func main() {
	// Thanks http://stackoverflow.com/a/3751456/1224232
	var minutes int

	// While structure in Go. While minutes is less than 1...
	for minutes < 1 {
		fmt.Printf("minutes: ")

		// We pass a pointer of our minutes variable for assignment
		_, err := fmt.Scanf("%d", &minutes)

		// If we have an error, which covers almost anything, we keep looping
		if err != nil {
			fmt.Println("We need a positive number")
		}
	}

	fmt.Printf("bottles: %d\n", minutesToWaterBottles(minutes))
}
