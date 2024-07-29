package main

import "fmt"

func closure1() func () int {
  counter := 0
  return func() int {
    counter = counter + 1
    return counter
  }
}

func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func filter(numbers []int, predicate func(int) bool) []int {
  final := []int {}
  fmt.Println("final ->", final)
  for _,v := range numbers {
    if predicate(v) {
      final = append(final, v)
      fmt.Println("final ->", final)
    }
  }
  return final
}

func main(){
  cl := closure1()
  for _ = range 5 {
    fmt.Printf(" %v ", cl())
  }
  fmt.Println()

  cl1 := adder()
  for i := range 10 {
    fmt.Printf(" %v", cl1(i))
  }
  fmt.Println()

  num := make([]int, 0)
  for i := range 10 {
    num = append(num, i)
  }
  
  even := filter(num, func(x int) bool {
    return x % 2 == 0
  })

  fmt.Println(even)
  
}
