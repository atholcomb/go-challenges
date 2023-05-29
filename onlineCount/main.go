// written by: atholcomb
// main.go
// returns the count of users online

package main

import "fmt"

func main() {
  fmt.Println(onlineCount(map[string]string{"Alice":"offline", "Bob":"online", "Jenny":"online"}))
}

func onlineCount(m map[string]string) []string {
  var result []string

  for k, v := range m {
    result = append(result, k+":"+v)
  }
  return result
}
