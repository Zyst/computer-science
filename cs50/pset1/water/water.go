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

  fmt.Printf("minutes: ")
  _, err := fmt.Scanf("%d", &minutes)

  // We should add error (Loop) handling for minutes input, make sure # is positive, and a number
  if err != nil {
    panic(err)
  }

  fmt.Printf("bottles: %d\n", minutesToWaterBottles(minutes))
}
