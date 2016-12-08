package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Our first test, we are only checking
// American Express, Mastercard, and Visa.
// We check that the passed card has a length of
// 13 characters, 15 characters, or 16 characters.
func checkCardLengthValidity(card string) bool {
	// If our length is 13, 15, or 16 we have passed the first test
	return len(card) == 13 || len(card) == 15 || len(card) == 16
}

// We make sure everything we input is a number.
// If someone types 378282246310005abcd that's invalid.
func checkInputIsNumbers(card string) bool {
	_, err := strconv.Atoi(card)

	// If we have no trouble converting Atoi we should be safe.
	return err == nil
}

func verifyCardWithLuhnAlgorithm(card string) bool {
	// We split our card into individual numbers
	numbers := strings.Split(card, "")

	// Our first sum following the algorithm
	firstSum := 0

	// Our second sum following the algorithm
	secondSum := 0

	for index, num := range numbers {
		// If it's an alternating number
		if index%2 != 0 {
			i, err := strconv.Atoi(num)

			if err != nil {
				panic(err)
			}

			sum := i * 2

			// If sum is less than 10
			if sum < 10 {
				firstSum += sum
			} else {
				// We turn our multiplied sum back to a string
				sumString := strconv.Itoa(sum)

				// We split the string into parts (Should just be two)
				split := strings.Split(sumString, "")

				// We convert the first element back to an int
				x, err := strconv.Atoi(split[0])

				if err != nil {
					panic(err)
				}

				// We convert our second element back to an int
				y, err := strconv.Atoi(split[1])

				if err != nil {
					panic(err)
				}

				// We add X and Y to our first sum
				firstSum = firstSum + x + y
			}
		} else {
			i, err := strconv.Atoi(num)

			if err != nil {
				panic(err)
			}

			secondSum += i
		}
	}

	// We convert our first sum and second sum into a total
	total := strconv.Itoa(firstSum + secondSum)

	// We break it into a slice
	totalSlice := strings.Split(total, "")

	// If our last number is 0, this card is "valid"
	return totalSlice[len(totalSlice)-1] == "0"
}

// Mastercard, Amex, and Visa are valid in our program
func returnCardType(card string) string {
	// We get the first two characters, which are the relevant ones to see types
	firstTwo := strings.Split(card, "")

	if firstTwo[0] == "5" &&
		firstTwo[1] == "1" ||
		firstTwo[1] == "2" ||
		firstTwo[1] == "3" ||
		firstTwo[1] == "4" ||
		firstTwo[1] == "5" {
		return "Mastercard"
	} else if firstTwo[0] == "4" {
		return "Visa"
	} else if firstTwo[0] == "3" &&
		firstTwo[1] == "4" ||
		firstTwo[1] == "7" {
		return "American Express"
	} else {
		return "Invalid"
	}
}

// This just combines all the functions above
func validateCardAndReturnType(card string) string {
  if checkCardLengthValidity(card) && checkInputIsNumbers(card) && verifyCardWithLuhnAlgorithm(card) {
    return returnCardType(card)
  } else {
    return "Invalid"
  }
}

func main() {
	// Here are some credit cards for testing!
	// https://www.paypalobjects.com/en_US/vhelp/paypalmanager_help/credit_card_numbers.htm
	test := "378282246310005"

  fmt.Println(validateCardAndReturnType(test))
}
