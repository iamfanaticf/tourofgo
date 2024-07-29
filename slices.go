package main

import (
  "fmt"
)

func main() {
  primes := [6]int {2, 3, 5, 7, 11, 13}

  var slice []int = primes[1:4]
  fmt.Println(slice)

  // slices are references to array. they do not store data
  names := [4]string {
    "Jhon",
    "Paul",
    "George",
    "Ringo",
  }
  fmt.Println(names)
  fmt.Println()

  a := names[0:2]
  b := names[1:3]
  fmt.Println(a,b)
  fmt.Println()
  b[0] = "changed_value"
  fmt.Println(a, b)
  fmt.Println()
  fmt.Println(names)

  fmt.Println()
  
  b = append(b, "abc", "def", "ghi")
  fmt.Println("slice b now is ",b)

  fmt.Println(names)
  // now changing the value of b will not reflect as underlying arr ref is changed
  b[0] = "new_changed_value"
  fmt.Println(names)
  fmt.Println()
  // slice literal is same as array literal without length
  sl := []struct {
    i int
    b bool
  } {
    {2,true},
    {3,false},
  }
  fmt.Println(sl)

  // slice defalut
  fmt.Println()
  fmt.Println(slice)
  fmt.Println(slice[1:])
  fmt.Println(slice[0:])
  fmt.Println(slice[:1])
  fmt.Println(slice[:])

  // slice length and capacity
  printSlice(slice)
  slice = slice[:2]
  printSlice(slice)
  slice = slice[:5]
  printSlice(slice)

  var nilSlice []int
  fmt.Println("  is slice nil", nilSlice == nil)

  sm := make([]int, 5, 10)
  fmt.Printf("%v is of length %v and cap %v \n", sm, len(sm), cap(sm))

  sos := [][]int {
    []int {1,2,3,4},
    []int {5,6,7,8},
  }
  fmt.Println(sos)
  // appending to a slice
  sos = append(sos, []int{9,9,9,9})
  fmt.Println(sos)
  sos = append(sos, []int{8,8,8})
  fmt.Println(sos)
  fmt.Println()
  for _, v := range sos {
    for _, v := range v {
      fmt.Printf(" %v ", v)
    }
    fmt.Println()
  }
}

func printSlice(s []int) {
  fmt.Println()
  fmt.Printf("\t %v  length: %v and capacity: %v", s, len(s), cap(s))
  fmt.Println()
}
