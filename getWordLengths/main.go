// written by: atholcomb
// main.go
// Given a word, stores the length of each word
// and returns a slice with the word and length

package main

import (
    "fmt"
    "strconv"
)

func main() {
  fmt.Println(printLengths(&[]string{"cat", "dog", "giraffe", "penguin", "wolf", "liger", "puma"}))
}

func printLengths(words *[]string) []string {
  var lengths []string

  for _, w := range *words {
    lengths = append(lengths, w+": "+strconv.Itoa(len(w)))
  }

 /* returns actual slice of values */
 return lengths
}
