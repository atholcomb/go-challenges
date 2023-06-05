// written by: atholcomb
// main.go
// FizzBuzz Interview question
// Write a function (fizzBuzz), that
// returns FizzBuzz, Fizz, or Buzz
// if a number is divisible by 3 or 5

package main

import ( 
    "fmt"
)

func main() {
  fmt.Println(fizzBuzz(3))
  fmt.Println(fizzBuzz(5))
  fmt.Println(fizzBuzz(15))
  fmt.Println(fizzBuzz(16))
}

func fizzBuzz(num int) (string, int){
  if num % 3 == 0 && num % 5 == 0 {
    return "FizzBuzz:", num
  } else if num % 3 == 0 {
    return "Fizz:", num
  } else if num % 5 == 0 {
    return "Buzz:", num
  } else {
    return "3 or 5 are not divisible into:", num
  }
}
