// written by: atholcomb
// main.go
// Given an original string, encrypt string with vowels
// replaced with key index value + aca added to the end

package main

import (
  "fmt"
  "strings"
)

func main() {
  /* setup header row to identify original string vs encrypted string */
  fmt.Printf("%s\t%s\n", "string", "eString")

  fmt.Println(karacaAlgo("apple"))
  fmt.Println(karacaAlgo("banana"))
  fmt.Println(karacaAlgo("karaca"))
  fmt.Println(karacaAlgo("burak"))
  fmt.Println(karacaAlgo("alpaca"))
}

func karacaAlgo(str string) string {
  var result string
  var encryptedStr string
  key := map[string]string{
    "a":"0",
    "e":"1",
    "i":"2",
    "o":"2",
    "u":"3",
  }
  
  /* reverse the original given string */
  for i := (len(str)-1); i >= 0; i-- {
    encryptedStr += string(str[i])
  }
  
  /* loop over key + encrypted string and replace vowels with key index value */
  for k, v := range key {
    for _, s := range encryptedStr {
      if k == string(s) {
        encryptedStr = strings.Replace(encryptedStr, string(s), v, 1)
      }
    }
  }

  /* format encrypted string with tab spacing and add "aca" at end of string */  
  result = fmt.Sprintf("%s\t%s%s", str, encryptedStr, "aca")
  return result
}
