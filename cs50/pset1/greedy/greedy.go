package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Recursive function that gives the coint count for the money
// you input following a Greedy algorithm. It will always try to give
// the least amount of coins possible.
func giveChange(money float64, coins int) int {
	// Coins denomination used are USD, so:
	// Quarters (25¢)
	// Dimes (10¢)
	// Nickels (5¢)
	// Pennies (1¢)

  // fmt.Printf("Money %f, Coins %d\n", money, coins)

	if money >= .25 {
		return giveChange(money-.25, coins+1)
	} else if money >= .10 {
		return giveChange(money-.10, coins+1)
	} else if money >= .05 {
		return giveChange(money-.05, coins+1)
	} else if money >= .01 {
    fmt.Println("Yay 41 came here")
		return giveChange(money-.01, coins+1)
	} else {
		// We have no change left to give!
		// Return coin count
		return coins
	}
}

func main() {
	var money float64

	// We greet people at first, because we are super nice.
	fmt.Printf("Oh hai! ")

	// While our heimoneyght is less than 1c
	for money < 0.01 {
		fmt.Printf("How much change is owed? ")

		// We grab a new reader and feed it the OS STDIN
		reader := bufio.NewReader(os.Stdin)

		// We read the input from our terminal
		input, _, err := reader.ReadLine()

		if err != nil {
			panic(err)
		}

		// We convert our []byte into a string, and then into an float, with
		// 64 bits, so float64
		money, err = strconv.ParseFloat(string(input), 64)
	}

	fmt.Println(giveChange(money, 0))
}
