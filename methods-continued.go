package main

import (
  "fmt"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func main() {
  f := MyFloat(10)
  f2 := MyFloat(-10)
  fmt.Println(f.Abs())
  fmt.Println(f2.Abs())
}
