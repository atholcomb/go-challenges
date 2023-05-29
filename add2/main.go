// written by: atholcomb
// main.go
// Adds two numbers and outputs the sum

package main

import ( 
    "fmt"
)

func main() {
  var a_num int
  var b_num int

  fmt.Printf("Enter value 1: ")
  fmt.Scanf("%v", &a_num)
  fmt.Printf("Enter value 2: ")
  fmt.Scanf("%v", &b_num)

  fmt.Printf("Sum of %v + %v is: %v\n", a_num, b_num, a_num + b_num)
}
