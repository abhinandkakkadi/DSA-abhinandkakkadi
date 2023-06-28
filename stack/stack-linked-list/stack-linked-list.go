package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type stack struct {
	head *Node
	size int
}

func main() {

	s := stack{nil, 0}

	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(23)
	s.print()
	fmt.Println()
	fmt.Println(s.Size())
	s.isEmpty()
	s.peek()
	fmt.Println()
	s.pop()
	s.print()
}

// push value into the stack
func (s *stack) push(value int) {

	node := &Node{value, nil}

	if s.head == nil {
		s.head = node
		s.size++

	} else {
		node.next = s.head
		s.head = node
		s.size++
	}
}

// pop value from the stack
func (s *stack) pop() {

	value := s.head.data
	s.head = s.head.next
	fmt.Println("popped element is ", value)
}

// print values present in the stack
func (s *stack) print() {

	temp := s.head
	if temp == nil {
		fmt.Println("stack empty")
		return
	}

	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

// return the size of the stack
func (s *stack) Size() int {

	return s.size
}

// check whether the stack is empty
func (s *stack) isEmpty() {

	if s.head == nil {
		fmt.Println("stack empty")
	} else {
		fmt.Println("stack not empty")
	}
}

// find the top element of the stack
func (s *stack) peek() {

	if s.head == nil {
		fmt.Println("stack empty")
	} else {
		fmt.Println(s.head.data)
	}
}
