// written by: atholcomb
// main.go
// Iterate through the alphabet
// and output the rune literal of each 
// letter as it's Unicode representation

package main

import ( 
    "fmt"
)

func main() {
  letters := "abcdefghijklmnopqrstuvwxyz"

  fmt.Println("Rune\tString")
  for _, l := range letters {
    fmt.Printf("%v\t%c\n", l, l)
  }
}
