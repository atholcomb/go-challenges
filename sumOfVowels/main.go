// written by: atholcomb
// main.go
// Create a function that takes a string and 
// returns the sum of vowels where some vowels 
// are considered numbers.
// Vowel	Number
// A	    4
// E	    3
// I	    1
// O	    0
// U	    0

package main

import (
  "fmt"
)

func main() {
  fmt.Println(sumOfVowels("Let's test this function."))
  fmt.Println(sumOfVowels("Do I get the correct output?"))
  fmt.Println(sumOfVowels("I love edabit!"))
}

func sumOfVowels(str string) int {
  var count int
  var vowelMap = map[string]int{
    "a":4,
    "e":3,
    "i":1,
    "o":0,
    "u":0, 
    "A":4,
    "E":3,
    "I":1,
    "O":0,
    "U":0, 
  }
  
  for k, v := range vowelMap {
    for _, s := range str {
      if string(s) == k {
        count = count + v
      }
    } 
  }
  return count
}
