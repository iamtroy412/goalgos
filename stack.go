package main

import (
	"fmt"
	"math"
)

type SNode struct {
    val interface{}
    prev *SNode
}

type Stack struct {
   Length int
   head *SNode
}

func (s *Stack) peek() interface{} {
    return s.head.val
}

func (s *Stack) push(d interface{}) {
    node := &SNode{val: d}
    s.Length++
    if s.Length == 1 {
        s.head = node
        return
    }
    
    node.prev = s.head
    s.head = node
}

func (s *Stack) pop() interface{} {
   s.Length = int(math.Max(0.0, float64(s.Length - 1))) 
   if s.Length == 0 {
       head := s.head
       s.head = nil
       return head.val
   }

   head := s.head
   s.head = s.head.prev
   return head.val
}

func main() {
    s := Stack{}
    s.push(4)
    s.push(8)
    s.push(88)
    s.push(42)
    s.push(69)

    fmt.Println("Length", s.Length)
    fmt.Println(s.peek())
    for i := 0; i <= s.Length; i++ {
        fmt.Println(s.pop())
    }

}
