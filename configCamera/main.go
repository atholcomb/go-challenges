// written by: atholcomb
// main.go
// Mock camera configuration 
// Sets frame rate value for each camera

package main

import ( 
    "fmt"
)

func main() {
  fmt.Println(configCamera())
}

func configCamera() map[string]int {
  var frame_rates []int
  var camera_frames int
  camera_defaults := map[string]int {
    "camera 1": 0, 
    "camera 2": 1, 
    "camera 3": 2, 
  }
  
  for k, _ := range camera_defaults {
    fmt.Printf("Enter camera %s frame rate: ", k)
    fmt.Scanf("%d", &camera_frames)   
    frame_rates = append(frame_rates, camera_frames)
  }
  for k, v := range camera_defaults {
    camera_defaults[k] = frame_rates[v]
  }
  fmt.Println()
  return camera_defaults
}
