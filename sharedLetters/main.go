package main

import (
    "fmt"
)

func main() {
  fmt.Println(sharedLetters("apple", "meaty"))
  fmt.Println(sharedLetters("fan", "forsook"))
  fmt.Println(sharedLetters("spout", "shout"))
}

func sharedLetters(s1, s2 string) (string, []string) {
  var letters []string
  var result string

  for _, i := range s1 {
    for _, j := range s2 {
      if i == j {
        letters = append(letters, string(i))
        result = s1 + " and " + s2 + " contain"
      }
    }
  } 
  return result, letters
}
