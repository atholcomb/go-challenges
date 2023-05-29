// written by: atholcomb
// main.go
// Checks given string for double letters

package main

import (
    "fmt"
    "strings"
)

func main() {
  ChkForDblLetters("yell")
  ChkForDblLetters("nnnn")
  ChkForDblLetters("aa")
  ChkForDblLetters("nono")
  ChkForDblLetters("yellow")
  ChkForDblLetters("ppeter")
}

func ChkForDblLetters(word string) {
  var count int
  var dblLetters []string
  letters := "abcdefghijklmnopqrstuvwxyz"
  for _, dl := range letters {
    dblLetters = append(dblLetters, string(dl)+string(dl))
  } 

  for _, dl := range dblLetters {
    if strings.Count(word, dl) >= 1 {
      count = strings.Count(word, dl)
      fmt.Println(word, count)
    }
  }
}
