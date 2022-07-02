package stacks

import "fmt"

type Node struct {
	data int
	next *Node
}

type StackLinkedList struct {
	top  *Node
	size int
}

func (s *StackLinkedList) push(val int) {
	// it is backwards
	s.top = &Node{val, s.top}
	s.size++
}

func (s *StackLinkedList) pop() int {
	if s.top == nil {
		fmt.Println("Stack underflow")
		return 0
	}
	popped := s.top.data
	s.top = s.top.next
	s.size--
	return popped
}

func (s *StackLinkedList) len() int {
	return s.size
}

func (s *StackLinkedList) isEmpty() bool {
	return s.size == 0
}

func (s *StackLinkedList) peek() int {
	return s.top.data
}

func (s *StackLinkedList) print() {
	curr := s.top
	fmt.Println("Printing the stack from top to bottom.")
	for curr != nil {
		fmt.Println(curr.data)
		curr = curr.next
	}
	fmt.Print()
}

func StackWithLinkedList() {
	var stack StackLinkedList
	stack.push(1)
	stack.push(2)
	stack.push(3)

	stack.print()

	stack.pop()
	fmt.Println("Top of stack is: ", stack.peek())
	stack.print()

	fmt.Println("Size of stack is :", stack.len())
	fmt.Println("Is stack empty: ", stack.isEmpty())
}
