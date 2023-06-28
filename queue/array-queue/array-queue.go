package main

import "fmt"

type Queue struct {
	data []int
	head int
	tail int
}

func main() {

	q := Queue{make([]int, 0), -1, -1}
	q.Add(2)
	q.Add(3)
	q.Add(4)
	q.Add(5)
	q.Add(5)
	q.Add(7)
	q.print()
	fmt.Println()

	q.Remove()
	q.print()
	fmt.Println()

	q.Remove()
	q.print()
	fmt.Println()

	q.Remove()
	q.print()
	fmt.Println()

	q.Remove()
	q.print()
	fmt.Println()

	q.Remove()
	q.print()
	fmt.Println()

	q.Remove()
	q.print()
	fmt.Println()

	q.Remove()

}

func (q *Queue) Add(value int) {

	if q.head == -1 {
		q.head = 0
		q.tail = 0

	} else {
		q.tail++
	}
	q.data = append(q.data, value)
}

func (q *Queue) Remove() {

	if q.head == -1 {
		fmt.Println("empty queue")
		return
	}
	q.head++
	if q.head == q.tail+1 {
		q.head = -1
		q.tail = -1
	}

}

func (q *Queue) isEmpty() {

	if q.head == -1 {
		fmt.Println("empty")
	}
}

func (q *Queue) print() {

	temp := q.head

	if temp == -1 {
		fmt.Println("queue empty")
		return
	}

	for temp != q.tail+1 {
		fmt.Println(q.data[temp])
		temp++
	}
}
