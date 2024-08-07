package main

import (
  "fmt"
  "math"
)

type MyFloat float64

type Vertex struct {
  X, Y float64
}

type Abser interface {
  Abs() float64
}

func (f MyFloat) Abs() float64 {
  if f < 0 {return float64(-f)}
  return float64(f)
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
} 

func main() {
  var a Abser 
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3, 4}

  a = f // since f implements the methods 
//a = v // will cause error bcz Abs is not been implemented on v 
  a = &v // this works as *Vertex is implemented
  fmt.Println(a.Abs())
}
