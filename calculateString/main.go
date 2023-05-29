// written by: atholcomb
// main.go
// Calculates the length of a string

package main

import (
    "fmt"
)

func main() {
  var value string

  fmt.Printf("Please enter a string: ")
  fmt.Scanf("%s", &value)
  fmt.Printf("Length of string '%s': %v\n", value, len(value))
}
