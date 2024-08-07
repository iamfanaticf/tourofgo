package main

import "fmt"

type I interface {
  M() 
}

type T struct {
  S string
}

type F float64

func (t T) M() {
  fmt.Println(t.S)
}

func (f F) M() {
  fmt.Println(f)
}

func main() {
  var i I
  i = T {"i am a struct"}
  describe(i)
  i.M()
  i = F(0.01)
  describe(i)
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T) ", i, i)
}
