// written by: atholcomb
// main.go
// Test program to try things out

package main

import (
  "fmt"
)

func main() {
  fmt.Println(sumOfOddEven(
    []int{1,2,3,4,5,6},
    []int{-1,-2,-3,-4,-5,-6}))
}

func sumOfOddEven(n []int, m []int) ([]int, []int) {
  var posSum []int
  var negSum []int
  var posEven int
  var posOdd int
  var negEven int
  var negOdd int

  /* sum even/odd positive numbers */
  for i := 1; i < (len(n)+1); i++ {
    if i % 2 == 0 {
      posEven += i
    } else {
      posOdd += i 
    }
  }
  posSum = append(posSum, posEven, posOdd)

  ///* sum even/odd negative numbers */
  for i := -1; i > (-len(m)-1); i-- {
    if i % 2 == 0 {
      negEven += i
    } else { 
      negOdd += i
    }
  }

  /* return positive/negative list */
  //return posSum, negSum
  negSum = append(negSum, negEven, negOdd)
  return posSum, negSum
}
