package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var height int

	// While our height is less than 1, more than or equal to 23
	for height < 1 || height > 23 {
		fmt.Printf("Height: ")

		// We grab a new reader and feed it the OS STDIN
		reader := bufio.NewReader(os.Stdin)

		// We read the input from our terminal
		input, _, err := reader.ReadLine()

		if err != nil {
			panic(err)
		}

		// We convert our []byte into a string, and then into an int, base10
		height, err = strconv.Atoi(string(input))

		// If we have an error, a negative number, or 0 we keep looping
		if err != nil || height < 1 {
			fmt.Println("We need a positive number")
		}
	}

  for index := 1; index <= height; index++ {
    spaces := strings.Repeat(" ", height - index)

    blocks := strings.Repeat("#", index) + "#\n"

    fmt.Printf(spaces + blocks)
  }
}
