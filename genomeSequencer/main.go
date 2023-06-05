// written by: atholcomb
// main.go
// Program creates sequences of dna strands

package main

import ( 
    "fmt"
    "math/rand"
    "time"
)

func main() {
  // # of sequences to be returned, but not including 21
  genomeSequencer(21)
}

func genomeSequencer(count int) {
  var genome = []string{"a", "t", "g", "c"}

  fmt.Println("GENOME SEQUENCER | SEQUENCE STARTS IN 3 SECONDS")
  fmt.Println("-----------------------------------------------")
  // postpone sequence start by 3 seconds
  time.Sleep(3 * time.Second)     

  for i := 1; i < count; i++ {
    // sleep 1 second before generating new sequence
    time.Sleep(1 * time.Second)   
    fmt.Printf("SEQUENCE:%02d ", i)
    for j := 0; j < 15; j++ {
      // generate a random index to be used in selecting genome letter
      var choice = rand.Intn(4)   
      fmt.Printf("%s", genome[choice])
    }
  fmt.Printf("%s", " -> VALIDATED SEQUENCE")
  fmt.Println()
  }
}
