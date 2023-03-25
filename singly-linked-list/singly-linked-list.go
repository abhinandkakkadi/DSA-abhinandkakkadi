package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func main() {

	list := LinkedList{}

	// Add Node at the beginning
	// list.addAtHead(1)
	// list.addAtHead(2)
	// list.addAtHead(3)
	// list.addAtHead(4)
	list.displayElements()
	fmt.Println()

	// Add Node at the end

	list.addAtEnd(1)
	list.addAtEnd(1)
	list.addAtEnd(1)
	list.addAtEnd(2)
	list.addAtEnd(2)
	list.addAtEnd(3)
	list.addAtEnd(3)
	list.addAtEnd(3)
	list.addAtEnd(3)
	list.displayElements()
	fmt.Println()

	//  Delete Node with value specified

	list.deleteValue(4)
	list.displayElements()
	fmt.Println()


	// 	Insert Nodes before and after a specified node

	list.addNodeBeforeAfter(7)
	list.displayElements()
	fmt.Println()

	// Delete duplicates in sorted Linked List

	list.removeDuplicates()
	list.displayElements()
	fmt.Println()

}





// Add elements at the head of Linked List

func (list *LinkedList) addAtHead(value int) {

	node := &Node{value,nil}

	if list.head == nil {
		list.head = node
	} else {
		node.next = list.head
		list.head = node
	}
}

// Add elements at the end of the linked list

func (list *LinkedList) addAtEnd(value int) {

	node := &Node{value,nil}
	temp := list.head

	if temp == nil {
		list.head = node
	} else {
		for temp.next != nil {
			temp =temp.next
		}
		temp.next = node
	}
}


// Delete a Node with value specified

func (list *LinkedList) deleteValue(value int) {
	temp := list.head

	if temp == nil {
		fmt.Println("List empty")
		return
	}

	if temp.data == value {
		list.head = temp.next
		return
	}

	for temp.next != nil {
		if temp.next.data == value {
			temp.next = temp.next.next
			return
		}
		if temp.next != nil {
			temp = temp.next
		}

	}
}

// Add Node Before And After 

func (list *LinkedList) addNodeBeforeAfter(value int) {

	temp := list.head 
	node1 := &Node{5,nil}
	node2 := &Node{5,nil}

	if temp == nil {
		fmt.Println("Empty list")
		return
	}

	if temp.data == value {
		node2.next = temp.next
		temp.next = node2

		node1.next = temp
		list.head = node1
		return
	}

	for temp.next != nil {

		if temp.next.data == value {

			node2.next = temp.next.next
			temp.next.next = node2

			node1.next = temp.next
			temp.next = node1
			return
		}
		if temp.next != nil {
			temp = temp.next
		}
		
	}


}

// Remove duplicates from the linked list in Sorted Linked list


func (list *LinkedList) removeDuplicates() {

	temp := list.head

	for temp.next!= nil {
		for temp.data == temp.next.data {
			temp.next =temp.next.next

			if temp.next == nil {
				return
			}
		}

		if temp.next != nil {
			temp = temp.next
		}
	}
}



// display elements of the linked list 

func (list *LinkedList) displayElements() {

	temp := list.head

	if temp == nil {
		fmt.Println("List empty")
	} else {
		for temp!= nil {
			fmt.Println(temp.data)
			temp= temp.next
		}
	}
}

