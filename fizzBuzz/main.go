// written by: atholcomb
// main.go
// Traditional fizzbuzz program

package main

import (
    "fmt"
    "strconv"
)

func main() {
  fmt.Println(fizzbuzz(3))
  fmt.Println(fizzbuzz(5))
  fmt.Println(fizzbuzz(15))
  fmt.Println(fizzbuzz(4))
}

func fizzbuzz(n int) (int string) {
  var value string

  if n % 3 == 0 && n % 5 == 0 {
    value = " FizzBuzz"
  } else if n % 3 == 0 {
    value = " Fizz" 
  } else if n % 5 == 0 {
    value = " Buzz"
  } else {
    value = " not divisible by 3 or 5"
  }
  return strconv.Itoa(n) + value
}
