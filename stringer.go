package main

import "fmt"

type Person struct {
  Name string
  Age int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
  a := Person{"Arthur Dent", 1}
  z := Person{"Zaphod Beeblebrox", 0}
  fmt.Println(a, z)
}
