// written by: atholcomb
// main.go
// Test struct called Vape

package main

import (
    "fmt"
)

type vape struct {
  modName string
  power int
  podSize float64
  ohmStrength float64
  unitType string 
}

func main() {
  device := vape{
    modName: "Caliburn A2",
    power: 1,
    podSize: 24.0,
    ohmStrength: 0.8,
    unitType: "Pod System",
  }
  
  fmt.Println(device.ohmStrength)
  device.updateOhm(0.9)
  device.printModName()
  fmt.Println(device.ohmStrength)
  
}

/* print all attributes from the struct */
func (v vape) print() {
  fmt.Println(v)
}

/* print only one attribute from the struct */
func (v vape) printModName() {
  fmt.Println(v.modName)
}

/* Update the Ohm value for Vape */
func (v *vape) updateOhm(newOhm float64) {
  (*v).ohmStrength = newOhm
}
