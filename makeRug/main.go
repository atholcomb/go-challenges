// written by: atholcomb
// main.go
// Outputs a rug with n x n column and rows
// with a specified character given

package main

import (
  "fmt"
  "strings"
)

func main() {
  makeRug() 
}

func makeRug() {
  var symbol string
  var columns int
  var rows int

  fmt.Printf("Enter character: ")
  fmt.Scanf("%v", &symbol)

  fmt.Printf("Enter how many columns: ")
  fmt.Scanf("%v", &columns)

  fmt.Printf("Enter how many rows: ")
  fmt.Scanf("%v", &rows)

  for i := 0; i < rows; i++ {
    fmt.Println(strings.Repeat(symbol, columns))
  }
}
