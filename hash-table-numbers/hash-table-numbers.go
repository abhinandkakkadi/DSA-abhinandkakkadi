package main

import (
	"fmt"
)

const size = 7

type HashTable struct {
	array [size]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key int
	next *bucketNode
}




func (h *HashTable) Insert(value int) {

	index := hash(value)
	h.array[index].insert(value)
}

func (b *bucket) insert(value int) {

	node := &bucketNode{value,nil}
	if b.head == nil {
		b.head = node
		return
	} else {
		node.next = b.head
		b.head = node
	}
}

func (h *HashTable) Search(value int) bool {
	index := hash(value)
	return h.array[index].search(value)
}

func (b *bucket) search(value int) bool {
	
	temp := b.head

	if b.head == nil {
		return false
	}

	if b.head.key == value {
		return true
	}

	for temp != nil {
		if temp.key == value {
			return true
		}
		temp = temp.next
	}

	return false

}

func (h *HashTable) Delete(value int) {
	
	index := hash(value)
	h.array[index].delete(value)

}

func (b *bucket) delete(value int) {

	temp := b.head

	if temp == nil {
		return
	}

	if b.head.key == value {
		b.head = b.head.next
		return
	}

	for temp.next != nil {

		if temp.next.key == value {
			temp.next = temp.next.next
			return
		}
	}
}

func hash(value int) int {

	return value%size
}



func Init() *HashTable {

	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}



func main() {

	hashTable := Init() 

	hashTable.Insert(1)
	hashTable.Insert(8)

	fmt.Println(hashTable.Search(8))

	hashTable.Delete(8)
	fmt.Println(hashTable.Search(8))

}