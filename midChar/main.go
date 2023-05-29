// written by: atholcomb
// main.go
// returns the middle char of given string

package main

import (
    "fmt"
)

func main() {
  findMidChar("")
  findMidChar("abc")
  findMidChar("abcd")
  findMidChar("abcde")
  findMidChar("aaaa")
  findMidChar("abdjabc")
}

func findMidChar(str string) {
  if len(str) % 2 != 0 {
    fmt.Printf("'%s', middle char is: %s\n", str, string(str[len(str) / 2]))
  } else if str == "" {
    fmt.Println("empty string is invalid")
  } else {
    fmt.Printf("string %s is even\n", str)
  } 
}
