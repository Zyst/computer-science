package main

import (
	"fmt"
)

func minutesToWaterBottles(min int) int {
	return min * 192 / 16
}

func main() {
	// Thanks http://stackoverflow.com/a/3751456/1224232
	var minutes int

	for minutes < 1 {
		fmt.Printf("minutes: ")
		_, err := fmt.Scanf("%d", &minutes)
		if err != nil {
			fmt.Println("We need a positive number")
		}
	}

	fmt.Printf("bottles: %d\n", minutesToWaterBottles(minutes))
}
