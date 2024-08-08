package main

import "fmt"

func Index[T comparable] (s []T, x T) int {
  for i, v := range s {
    if v == x {
      return i
    }
  }
  return -1
}

func main() {
  intSlice := []int {1, 2, 4, 3, 5}
  stringSlice := []string {"one", "two", "four", "three", "five"}

  fmt.Println(Index(intSlice, 3))
  fmt.Println(Index(stringSlice, "three"))

}
