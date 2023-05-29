// written by: atholcomb
// main.go
// returns in a given pin is valid

package main

import (
    "fmt"
    "strings"
)

func main() {
  validPin("1234")
  validPin("45135")
  validPin("89abc1")
  validPin("900876")
  validPin(" 4983")
  validPin("00 00 00")
}

func validPin(pin string) {

  if len(pin) == 4 || len(pin) == 6 {
    if !strings.Contains(pin, "abc") {
      fmt.Printf("%v is a valid pin\n", pin)
    } 
  } else {
    fmt.Printf("%v is not a valid pin\n", pin)
  }
}
