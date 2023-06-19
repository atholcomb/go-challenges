// written by: atholcomb
// main.go
// Write a function (fizzBuzz), that
// returns FizzBuzz, Fizz, or Buzz
// if a number is divisible by 3, 5 or 15

package main

import (
  "fmt"
)

func main() {
  fmt.Println(fizzBuzz())
}

func fizzBuzz() (int, string) {
  var num int
  n, m, o := 3, 5, 15

  fmt.Printf("Please enter a number: ")
  fmt.Scanf("%v", &num)

  if num % o == 0 {
    return num, "Fizzbuzz"
  } else if num % n == 0 {
    return num, "Fizz"
  } else if num % m == 0 {
    return num, "Buzz"
  } else {
    return num, "is not divisble by 3, 5, or 15"
  }
}
