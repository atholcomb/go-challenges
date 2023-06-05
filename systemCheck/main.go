// written by: atholcomb
// main.go
// Program outputs a mock system check analysis

package main

import ( 
    "fmt"
    "time"
    "strings"
)

func main() {
  env := sysCheck()

  fmt.Printf("%s%s\n", "Connected to system: ", strings.Title(env["hostname"]))
  fmt.Println("//Reading attributes//")

  for k, v := range env {
    time.Sleep(1 * time.Second)
    fmt.Printf("%v: %v\n", k, v)
  }
}

func sysCheck() map[string]string {
  env := map[string]string {
    "os": "linux",
    "hostname": "lrs-12-penguin-pc",
    "cpu_cores": "12",
    "gpu_type": "cuda",
    "ram_size": "128gb",
    "kernel": "5.2.1",
    "vm": "yes",
  }
  return env
}
