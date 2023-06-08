// written by: atholcomb
// main.go
// returns the length of an integer given

package main

import (
  "fmt"
  "strconv"
)

func main() {
  fmt.Println(numLength(10))
  fmt.Println(numLength(5000))
  fmt.Println(numLength(0))
}

func numLength(n int) int {
  var count int
  
  for i := range strconv.Itoa(n) {
    if i > -1 {
      count++
    } 
  }
  return count
}
