// written by: atholcomb
// main.go
// Create a function that converts a date 
// formatted as MM/DD/YYYY to YYYYDDMM.

package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(formatDate("11/12/2019"))
  fmt.Println(formatDate("12/31/2019"))
  fmt.Println(formatDate("01/15/2019"))
}

func formatDate(s string) string {
  str := strings.Split(s, "/")
  var output []string
  var result string

  for i := (len(str)-1); i >= 0; i-- {
    output = append(output, str[i])
  }
  result = fmt.Sprintf("%s -> %s", s, strings.Join(output, ""))

  return result
}
