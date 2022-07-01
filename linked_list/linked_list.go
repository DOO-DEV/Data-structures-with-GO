package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

func (head *Node) PrintList() {
	for {
		data := head.data
		fmt.Println(data)
		if head.next == nil {
			break
		}
		head = head.next
	}
}

func NewNode(data int) *Node {
	return &Node{data, nil}
}

func NewNodeWithNext(data int, next *Node) *Node {
	return &Node{data, next}
}

func (head *Node) AddAfter(data int, before *Node) {
	current := head
	for current != nil {
		if current.next == before {
			newNode := Node{data, before.next}
			before.next = &newNode
		}
		current = current.next

	}
}

func (head *Node) Append(value int) {
	curr := head
	if curr != nil {
		for curr.next != nil {
			curr = curr.next
		}
	}
	curr.next = NewNode(value)
}

func (head *Node) Unshift(data int) *Node {
	return NewNodeWithNext(data, head)
}

func (head *Node) Pop() {
	curr := head
	var before *Node
	if curr != nil {
		for curr.next != nil {
			before = curr
			curr = curr.next
		}
	}
	fmt.Println(before)
	before.next = nil

}

func (head *Node) Length() int {
	curr := head
	length := 0
	for curr != nil {
		length += 1
		curr = curr.next
	}

	return length
}

func RecursiveLength(node *Node) int {

	if node == nil {
		return 0
	}
	return 1 + RecursiveLength(node.next)
}

func LinkedList() {
	a := NewNode(1)
	b := NewNode(2)
	k := NewNode(9)
	c := NewNode(2)
	d := NewNode(2)

	a.next = b
	b.next = k
	k.next = c
	c.next = d

	a.AddAfter(10, c)
	fmt.Println(c.next.data)

	a.Append(12)
	fmt.Println(d.next.data)

	g := NewNode(12)
	p := NewNode(13)
	d.next = g
	g.next = p

	a.Pop()
	fmt.Println(g.next)

	newHead := a.Unshift(-1)
	fmt.Println(newHead.next)

	fmt.Println("length with iterative fucntiality: ", newHead.Length())
	newHead.PrintList()

	fmt.Println("lenght wiht recursive functionality: ", RecursiveLength(newHead))

}
