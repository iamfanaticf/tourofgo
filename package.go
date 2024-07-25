// package 
package main

// imports 
import(
  "fmt"
)

// function
func add(x,y int) int {
  return x + y
}

// multiple return values function
func swap(x, y int) (int, int) {
  return y, x
}

// named return values function
func swap2(x, y int) (a, b int) {
  a = y
  b = x
  return 
}

func main() {
  fmt.Println("hello world")
  fmt.Println("add func ->", add(1,2))
  a, b := swap(1, 2)
  fmt.Printf("swap -> %v, %v\n", a, b)
  a, b = swap2(1, 2)
  fmt.Println("swap2 -> ", a , b)


  
  // variable declaration
  var v1 int
  // default value
  fmt.Println(v1)
  v1 = 1
  fmt.Println(v1)
  fmt.Println()

  // var with initializers
  var v2 = 2
  fmt.Println(v2)
  
  // short variable declaration
  v3 := 3
  fmt.Println(v3)
  fmt.Println()

  // type conversions
  fmt.Printf("converting int into float32 type = %T",float32(1))

  // type inference 
  v4 := 2 + 3i
  fmt.Printf("Type of v4 is %T",v4)

  // const
  const v5 = 3.1415
  fmt.Println("const value", v5)

  // An untyped constant takes the type needed by its context.
  const v6 = 1 << 100
  fmt.Printf("val has type %T", v6)


}
