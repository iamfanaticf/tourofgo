package main

import (
  "fmt"
)

func compute(fn func(int, int)int, x, y int) int {
  return fn(x , y)
}

func main() {
  add := func(x,y int) int { return x + y}
  sub := func(x, y int) int { return x - y}
  mul := func(x, y int) int { return x * y}
  div := func(x, y int) int { return x/y}

  fmt.Println(compute(add, 3, 4))
  fmt.Println(compute(sub, 3, 4))
  fmt.Println(compute(mul, 3, 4))
  fmt.Println(compute(div, 3, 4))
}
