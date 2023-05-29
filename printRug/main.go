// written by: atholcomb
// main.go
// Outputs a repeated character n times

package main

import (
    "fmt"
    "strings"
)

func main() {
  var character string
  var count int

  fmt.Printf("Type a character: ")
  fmt.Scanf("%s", &character)

  fmt.Printf("Repeat how many times?: ")
  fmt.Scanf("%v", &count)

  for i := 0; i < count; i++ {
    fmt.Println(strings.Repeat(character, count))
  }
}
