package main

import "fmt"

func main() {
  // maps
  var m map[int]string
  fmt.Println("is map nil?", m == nil)

  v, ok := m[0]
  if(!ok) {
    fmt.Println("this is not ok")
  } else {
    fmt.Println(v)
  }

  m = make(map[int]string)
  m[1] = "1"
  fmt.Println("is map nil?", m == nil)

  delete(m, 1)
  elem, ok := m[1]
  if(!ok) {
    fmt.Println("this is not ok")
  } else {
    fmt.Println(elem)
  }

}
