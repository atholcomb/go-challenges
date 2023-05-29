// written by: atholcomb
// main.go
// returns a string n times

package main

import ( 
    "fmt"
    "strings"
)

func main() {
  fmt.Println(printLetters("s", 5))
}

func printLetters(l string, n int) string {
  var letters []string

  for i := 0; i < n; i++ {
    letters = append(letters, l) 
  }

  return strings.Join(letters, "")
}
