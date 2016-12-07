package main

import (
	"fmt"
)

// Our first test, we are only checking
// American Express, Mastercard, and Visa.
func checkCardLengthValidity(card string) bool {
  // If our length is 13, 15, or 16 we have passed the first test
  return len(card) == 13 || len(card) == 15 || len(card) == 16
}

func main() {
  // Here are some credit cards for testing!
  // https://www.paypalobjects.com/en_US/vhelp/paypalmanager_help/credit_card_numbers.htm

  // TODO: Add User input here
  test := "378282246310005"

  // If our card doesn't pass length tests we return
  if !checkCardLengthValidity(test) {
    fmt.Println("INVALID\n")
    return
  }
}
