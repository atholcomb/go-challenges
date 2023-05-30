// written by: atholcomb
// main.go
// returns a rune literal for a given rune

package main

import ( 
    "fmt"
)

func printRune(r rune) {
  fmt.Printf("%q  %d\n", r, r)
}

func main() {
  printRune('a')
  printRune('b')
  printRune('c')
}
