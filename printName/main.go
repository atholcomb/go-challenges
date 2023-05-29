// written by: atholcomb
// main.go
// Prints first and last name

package main

import ( 
    "fmt"
)

func main() {
  var firstname string
  var lastname string

  fmt.Printf("Please enter your first name: ")
  fmt.Scanf("%s", &firstname)

  fmt.Printf("Please enter your last name: ")
  fmt.Scanf("%s", &lastname)

  fmt.Println("Name entered:", firstname, lastname)
}

