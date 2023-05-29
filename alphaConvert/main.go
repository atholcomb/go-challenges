// written by: atholcomb
// main.go
// Iterate through a string of letters
// and output the index and each letter
// or output char, ANSI# and Unicode representation

package main

import ( 
    "fmt"
)

//func main() {
//  letters := "abcdefghi"
//  
//  fmt.Println("Char", "\tANSI#", "\tUnicode")
//  for i := 0; i < len(letters); i++ {
//    fmt.Printf("%c\t%v\t%U\n", letters[i], letters[i], letters[i])
//  }
//}

//func main() {
//  letters := "abcdefghi"
//
//  for i, letter := range letters {
//    fmt.Printf("%v %c\n", i, letter)
//  }
//}

func main() {
  letters := []string{"a","b", "c"}

  for i, letter := range letters {
    fmt.Println(i, letter)
  }
}
