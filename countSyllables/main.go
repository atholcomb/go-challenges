// written by: atholcomb
// main.go
// Counts the number of syllables
// for a given string

package main

import ( 
    "fmt"
    "strings"
)

func main() {
  countSyllables("ho-tel")
  countSyllables("cat")
  countSyllables("met-a-phor")
  countSyllables("ter-min-a-tor")
}

func countSyllables(str string) {
  var count int
  s := strings.ReplaceAll(str, "-", "")

  for _, vowel := range "aeiou" {
    if strings.Count(str, string(vowel)) >= 1 {
      count += 1
    } 
  }
  fmt.Println(s,":",count)
}
