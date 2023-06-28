package main

import "fmt"

type stack struct {
	data []int
	top  int
}

func main() {

	s := stack{make([]int, 0), -1}

	s.push(1)
	s.push(2)
	// s.push(3)
	// s.push(4)
	s.pop()
	fmt.Println()
	val := s.peek()
	if val != -1 {
		fmt.Println(val)
	}

	b := s.isEmpty()
	if b {
		fmt.Println("list empty")
	} else {
		fmt.Println("list not empty")
	}

}

func (s *stack) push(value int) {
	s.top++
	s.data = append(s.data, value)
}

func (s *stack) pop() {

	if s.top == -1 {
		fmt.Println("stack empty")
		return
	}
	s.top--
}

func (s *stack) peek() int {

	if s.top == -1 {
		fmt.Println("list empty")
		return -1
	} else {
		val := s.data[s.top]
		return val
	}
}

func (s *stack) isEmpty() bool {

	b := s.top
	return b == -1
}
