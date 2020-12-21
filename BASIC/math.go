package main

import (
  "fmt"
  "math"
)
func main() {
  var num float64 = 24
  var sqrtRes = math.Sqrt(num)
  fmt.Println("Square Root:", sqrtRes)
  fmt.Printf("The output is %g Thank you.", sqrtRes)
  fmt.Println()
  fmt.Printf("The output is %.2f Thank you.", sqrtRes)

  //ceil -> 3.1 = 4, floor -> 3.1 = 3
  var roundResult = math.Round(sqrtRes)
  fmt.Println()
  fmt.Println("Round Result:", roundResult)
  var ceilResult = math.Ceil(sqrtRes)
  fmt.Println("Ceil Result:", ceilResult)
  var floorResult = math.Floor(sqrtRes)
  fmt.Println("Floor Result:", floorResult)
}
