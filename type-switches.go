package main

import "fmt"

func do(i interface{}) {
  switch v := i.(type) {
  case int :
    fmt.Println("it is an int")
  case string :
    fmt.Println("it is string")
  default :
    fmt.Printf("I dunno %T\n", v)
  }
}

func main() {
  do(1)
  do("two")
}
