package main
import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v *Vertex) Scale(scale_factor float64) {
  v.X = v.X * scale_factor
  v.Y = v.Y * scale_factor
}

func main() {
  v := Vertex {3, 4}
  v.Scale(10)
  fmt.Println(v.Abs())
}

