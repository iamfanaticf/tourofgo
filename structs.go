package main

import "fmt"

type Vertex struct {
  X, Y int
}

func main() {
  fmt.Println(Vertex{1, 2})
  x := Vertex {2,3}
  fmt.Println("X is ", x)
  fmt.Println("X Vertex is", x.X)

  // pointer to struct
  p := &x
  (*p).Y = 9
  fmt.Println(x)
  p.X = 8
  fmt.Println(x)

  var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p2  = &Vertex{1, 2} // has type *Vertex
  )
  fmt.Println(v1, v2, v3, p2)
}
