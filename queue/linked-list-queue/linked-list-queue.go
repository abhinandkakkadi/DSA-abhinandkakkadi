package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type queue struct {
	head *Node
	tail *Node
	size int
}

func main() {

	l := queue{nil,nil,0}
	// adding elements to the queue
	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)
	l.Enqueue(4)
	l.print()
	fmt.Println()

	// removing elements from the queue
	l.Dequeue()
	l.print()
	fmt.Println()

	// peek the first element
	l.Peek()

	// check whether the queue is empty
	l.isNull()

	// returns the size of the queue 
	fmt.Println("size of the queue",l.Size())

}
// Add elements at the end
func (l *queue) Enqueue(value int) {

	node := &Node{value,nil}

	if l.head == nil {
		l.head = node
		l.tail = node
		l.size++
		return
	}

	l.tail.next = node
	l.tail  = node
	l.size++

}

// delete elements from the top
func (l *queue) Dequeue() {

	if l.head == nil {
		fmt.Println("queue empty")
		return
	}

	l.head = l.head.next
	l.size--

}

// Traversing through the queue
func (l *queue) print() {

	temp := l.head
	if temp == nil {
		fmt.Println("Empty queue")
		return
	}

	for temp != nil {
		fmt.Println(temp.data)
		temp=temp.next
	}
}

//  peek the first element
func (l *queue) Peek() {

	if l.head == nil {
		fmt.Println("queue empty")
		return
	}

	fmt.Println(l.head.data)

}

// check whether the queue is empty
func (l *queue) isNull() {

	if l.head == nil {
		fmt.Println("queue empty")
	} else {
		fmt.Println("queue not empty")
	}
}


// return the size of the queue
func (l *queue) Size() int {

	return l.size
}