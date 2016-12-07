package main

import (
	"fmt"
	"strconv"
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
  if err != nil {
    return false
  }
  return true
}

func main() {
  // Here are some credit cards for testing!
  // https://www.paypalobjects.com/en_US/vhelp/paypalmanager_help/credit_card_numbers.htm

  // TODO: Add User input here
  test := "378282246310005"

  // If our card doesn't pass length tests we return
  if !checkCardLengthValidity(test) && !checkInputIsNumbers(test) {
    fmt.Println("INVALID")
    return
  }
}
