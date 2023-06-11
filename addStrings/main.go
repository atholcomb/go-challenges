// written by: atholcomb
// main.go
// returns the sum of two strings 

package main

import (
  "fmt"
  "strconv"
)

func main() {
  fmt.Println(addStrings("111", "111"))
  fmt.Println(addStrings("10", "80"))
  fmt.Println(addStrings("", "20"))
}

func addStrings(n string, m string) string {
  n1, _ := strconv.Atoi(n) 
  n2, _ := strconv.Atoi(m)
  
  if n == "" || m == "" {
    return "Invalid Operation"
  }
  return strconv.Itoa(n1 + n2)
}
