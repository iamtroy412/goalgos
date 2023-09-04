package deque

type Node struct {
    val any
    next *Node
    prev *Node
}

type Deque struct {
    length int
    head *Node
    tail *Node
}

func (d *Deque) IsEmpty() bool {
    return d.length == 0
}

func (d *Deque) Size() int {
    return d.length
}

func (d *Deque) AddFist(item any) {
    newNode := &Node{
        val: item,
    }

    if d.head == nil {
        d.head = newNode
        d.tail = newNode
    } else {
        newNode.next = d.head
        d.head.prev = newNode
        d.head = newNode
    }

    d.length++
}

func (d *Deque) AddLast(item any) {
    newNode := &Node{
        val: item,
    }

    if d.head == nil {
        d.head = newNode
        d.tail = newNode
    } else {
        currentNode := d.head
        for currentNode.next != nil {
            currentNode = currentNode.next
        }
        newNode.prev = currentNode
        currentNode.next = newNode
        d.tail = newNode
    }
    d.length++
}
