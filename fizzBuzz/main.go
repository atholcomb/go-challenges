// written by: atholcomb
// main.go
// Write a function (fizzBuzz), that
// returns FizzBuzz, Fizz, or Buzz
// for all factors of 3, 5 or both 3 and 5

package main

import (
  "fmt"
  "strconv"
)

func main() {
  fizzBuzz()
}

func fizzBuzz() {
  var fbcount, fzcount, bzcount int
  var fbList, fzList, bzList []string
  var numList [31]int

  /* test for Fizzbuzz, Fizz and Buzz in range numList */
  for num := range numList { 
    if num % 3 == 0 && num % 5 == 0 {
      fbList = append(fbList, "Fizzbuzz: " + strconv.Itoa(num))
    } else if num % 3 == 0 {
      fzList = append(fzList, "Fizz: " + strconv.Itoa(num))
    } else if num % 5 == 0 {
      bzList = append(bzList, "Buzz: " + strconv.Itoa(num))
    }
  }

  /* output all factors divisble by 3 */
  fmt.Print("Divisble by 3:\n")
  for _, v := range fzList {
    fzcount++
    fmt.Printf("%v\n", v)
  }
  fmt.Printf("%s%d\n","Total count: ", fzcount)

  /* output all factors divisble by 5 */
  fmt.Print("\nDivisble by 5:\n")
  for _, v := range bzList {
    bzcount++
    fmt.Printf("%v\n", v)
  }
  fmt.Printf("%s%d\n","Total count: ", bzcount)

  /* output all factors divisble by both 3 and 5 */
  fmt.Print("\nDivisble by both 3 and 5:\n")
  for _, v := range fbList {
    fbcount++
    fmt.Printf("%v\n", v)
  }
  fmt.Printf("%s%d\n\n","Total count: ", fbcount)
}
