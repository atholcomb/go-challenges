// written by: atholcomb
// main.go
// returns the string and rune literal values together

package main

import ( 
    "fmt"
)

func stringToRune(s string) {
  for _, r := range s {
    fmt.Printf("%c  %d\n", r, r)
  }
  fmt.Println()
}

func main() {
  stringToRune("view")
  stringToRune("keg")
  stringToRune("regs")
}
