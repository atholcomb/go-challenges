// written by: atholcomb
// main.go
// sums 2 integers and returns
// true if the sum is less then 
// 100 and false otherwise

package main

import (
    "fmt"
)

func main() {
  fmt.Println(lessThan100(22, 15))
  fmt.Println(lessThan100(83, 34))
  fmt.Println(lessThan100(3, 77))

}

func lessThan100(n1, n2 int) bool {
  /* notice return statements included inside each condition block */

  if n1 + n2 < 100 {
    return true
  } else {
    return false
  }
}
