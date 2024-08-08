package main

import "fmt"

type List[T any] struct {
  val T
  next *List[T]
}
type LinkedList[T any] struct {
  head *List[T]
}
func (l *LinkedList[T]) addLast(val T) {
  head := l.head
  if head == nil {
    l.head = &List[T]{val,nil}
    return
  } 
  for head.next != nil {
    head = head.next
  } 
  head.next = &List[T]{val,nil}
}

func (l *LinkedList[T]) addFirst(val T){
  head := l.head
  if head == nil {
    l.head = &List[T]{val, nil}
  }
  newHead := &List[T]{val,nil}
  newHead.next = head
  l.head = newHead
}

func (l *LinkedList[T]) deleteLast(){
  head := l.head
  if head == nil {
    return
  }
  if head.next == nil {
    l.head = nil
    return
  } 
  for head.next.next != nil {
    head = head.next
  }
  head.next = nil
}

func (l *LinkedList[T]) deleteFirst() {
  head := l.head
  if head == nil {
    return
  }
  l.head = head.next
}

func main() {
  l := LinkedList[int]{&List[int]{1,nil}}
  l.addLast(2)
  fmt.Println(l.head.next.val)
}
