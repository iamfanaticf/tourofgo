package main

import(
  "fmt"
  "math/rand"
)

func main() {
  // basic for loop
  for i:=0; i<5; i+=1 {
    fmt.Printf(" %v ", i)
  }
  fmt.Println() 

  // go while = for 
  i := 1
  for i<5 {
    fmt.Printf(" %v ", i)
    i += 1
  }
  fmt.Println() 
  // infinite for 
  for {
    fmt.Printf(" this would iterate forever if not broken \n")
    if(i == 5) { break } 
  }
  
  // basic if 
  if i == 5 {
    fmt.Println(" i is 5 ")
  } else {
    fmt.Println(" i is not 5 ")
  }

  // if with a statement
  if v:=rand.Intn(10); v < 5 {
    fmt.Println(" v is less than 5")
  } else {
    fmt.Println(" v is greater than 5")
  }

  // switch 
  mycase := "ghi"
  switch mycase {
    case "abc" : fmt.Println(" case 1")
    case "def" : fmt.Println(" case 2")
    case "ghi" : fmt.Println(" case 3")
    default : fmt.Println(" no case matched")
  }

  // switch with no condition
  switch {
    case 1 > 2 : fmt.Println(" case 1")
    case 1 > 3 : fmt.Println(" case 2")
    default : fmt.Println(" no case matched")
  }

  // defer and defer stack
  defer fmt.Println(" three")
  defer fmt.Println(" two")
  defer fmt.Println(" one")

  defer fmt.Println(" this statement will print after below statement")
  fmt.Println(" i am below statement")




}
