package main

import (
  "fmt"
  "math/rand"
) 

func main() {
  var arr [5]int 
  for i := range 5 {
    arr[i] = rand.Intn(100)
  }
  fmt.Println(arr)

}
