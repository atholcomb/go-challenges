// written by: atholcomb
// main.go
// You are given a dictionary containing the cost price per unit (in dollars), 
// sell price per unit (in dollars), and the starting inventory. 
// Return the total profit made, rounded to the nearest dollar.

package main

import (
  "fmt"
)

func main() {
  fmt.Println(calcProfit(map[string]float64{
  "cost_price": 32.67,
  "sell_price": 45.00,
  "inventory": 1200.00,
  }))

  fmt.Println(calcProfit(map[string]float64{
  "cost_price": 225.89,
  "sell_price": 550.00,
  "inventory": 100.00,
  }))

  fmt.Println(calcProfit(map[string]float64{
  "cost_price": 2.77,
  "sell_price": 7.95,
  "inventory": 8500.00,
  }))
}

func calcProfit(p map[string]float64) string {
  var costTotal float64
  var sellDebit float64
  
  costTotal = p["cost_price"] * p["inventory"]
  sellDebit = p["sell_price"] * p["inventory"]
  
  return fmt.Sprintf("%.2f", sellDebit - costTotal)
}
