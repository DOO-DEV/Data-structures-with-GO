package stacks

import "fmt"

type Stack_array struct {
	maxSize int
	stack   []int
	top     int
}

func (s *Stack_array) push(val int) bool {
	if s.top >= s.maxSize-1 {
		fmt.Println("Stack overflow")
		return false
	}
	s.top++
	s.stack[s.top] = val
	fmt.Printf("%v is pushed into stack\n", val)
	return true
}

func (s *Stack_array) pop() int {
	if s.top < 0 {
		fmt.Println("Stack underflow")
		return -1
	}

	poped := s.stack[s.top]
	s.top--
	s.stack = s.stack[0 : s.top+1]
	return poped

}

func (s *Stack_array) isEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack_array) peek() int {
	fmt.Println("top is: ", s.top)
	return s.stack[s.top]
}

func (s *Stack_array) size() int {

	return len(s.stack)
}

func (s *Stack_array) print() {
	for _, val := range s.stack {
		fmt.Println(val)
	}
}

func StackWithArray() {
	// initialize stack
	stack := Stack_array{top: -1, maxSize: 5, stack: make([]int, 5)}

	stack.push(1)
	stack.push(2)
	stack.push(3)

	stack.print()

	fmt.Println("before pop, peek is: ", stack.peek())
	stack.pop()
	fmt.Println("printing the stackk: ")
	stack.print()
	fmt.Println("after pop, peek is: ", stack.peek())

}
