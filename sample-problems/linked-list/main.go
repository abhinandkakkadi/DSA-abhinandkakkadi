package main

import (
	"fmt"
)


type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {

	list := LinkedList{}
	list.addHead("abhinand")
	list.addHead("geetha")
	list.addHead("athira")

	// list.displayElements()
	// fmt.Println()

	// list.addAtEnd("ok fine")

	// list.displayElements()
	// fmt.Println()

	// list.deleteValue("geetha")

	// list.displayElements()
	// fmt.Println()

	list.addNodeBeforeAndAfter("abhinand")

	list.displayElements()
	fmt.Println()

	list.removeDuplicates()

}

func (l *LinkedList) addHead(val string) {

	l.head =  &Node{data: val,next: l.head}

}

func (l *LinkedList) displayElements() {

	temp := l.head

	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}

}

func (l *LinkedList) addAtEnd(val string) {

	temp := l.head
	node := &Node{data: val,next: nil}

	if temp == nil {
		l.head = node
	}

	for temp.next != nil {
		temp = temp.next
	}

	temp.next = node

}


func (l *LinkedList) deleteValue(val string) {

	temp := l.head

	if temp == nil {
		fmt.Println("list empty")
		return
	}

	if temp.data == val {
		l.head = l.head.next
		return
	}

	for temp.next != nil {
		if temp.next.data == val {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}

}

func (l *LinkedList) addNodeBeforeAndAfter(val string) {
	
	nodeBefore := &Node{data:"alwin",next: nil}
	nodeAfter := &Node{data: "unaiz",next: nil}
	temp := l.head

	if temp == nil {
		fmt.Println("list empty")
		return 
	}

	if temp.data == val {
		nodeAfter.next = temp.next
		temp.next = nodeAfter

		nodeBefore.next = temp
		l.head = nodeBefore
		return
	}

	for temp.next != nil {

		if temp.next.data == val {
			nodeAfter.next = temp.next.next
			temp.next.next = nodeAfter

			nodeBefore.next = temp.next
			temp.next = nodeBefore
		
			return
		}

		temp = temp.next
	}

}

