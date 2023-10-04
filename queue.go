package main

import "fmt"

type Node struct {
    value interface{}
    next *Node
    prev *Node
}

type Queue struct {
    Length int
    head *Node
    tail *Node
}

func (q *Queue) peek() interface{} {
    return q.head.value
}

func (q *Queue) deque() interface{} {
    if q.head == nil {
        return nil
    }

    q.Length--

    head := q.head
    q.head = q.head.next
    
    return head.value
}

func (q *Queue) enqueue(d interface{}) {
    node := &Node{value: d}
    q.Length++
    if q.tail == nil {
        q.tail = node
        q.head = node
        return
    }

    q.tail.next = node
    q.tail = node
}


func main() {
    q := Queue{}
    q.enqueue(4)
    q.enqueue(8)
    q.enqueue(88)
    q.enqueue(42)
    q.enqueue(69)

    fmt.Println("Length", q.Length)
    fmt.Println(q.peek())
    fmt.Println(q.deque())
    fmt.Println(q.deque())
    fmt.Println(q.deque())
    fmt.Println(q.deque())
    fmt.Println(q.deque())
}

