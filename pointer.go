package main

import(
  "fmt"
)

func main() {
  i := 1
  // pointer declaration
  var p *int
  // default value of pointer is nil
  fmt.Println("is p nil", p == nil)
  // assingment
  p = &i
  fmt.Println("address of var i is ", p)
  
  // dereferencing the pointer
  fmt.Println("underlying value of the pointer is", *p)

  // pointer to a pointer
  p2 := &p
  fmt.Printf("address of pointer is %v and address of pointer2 is %v \n", p, p2)

  // modifying the original variable through pointer
  **p2 = 2
  fmt.Println("the value of var i now is ", i)
  
}
