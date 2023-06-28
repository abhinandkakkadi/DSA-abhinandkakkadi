package main

import "fmt"

const size = 7

type HashTable struct {
	array [size]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func Init() *HashTable {

	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func main() {

	table := Init()

	// Insert elements into the Hash Table
	table.Insert("Abhinand")
	table.Insert("Athira")
	table.Insert("Geetha")

	// Search elements in the Hash Table
	fmt.Println(table.Search("Abhinand"))
	fmt.Println(table.Search("Athira"))

	// Delete elements in the Hash Table
	table.Delete("Abhinand")

	fmt.Println(table.Search("Abhinand"))
	fmt.Println(table.Search("Athira"))
	fmt.Println(table.Search("Geetha"))
}

// Insert elements into the Hash Table
func (h *HashTable) Insert(k string) {

	index := hash(k)
	h.array[index].insert(k)

}

func (b *bucket) insert(k string) {

	node := &bucketNode{k, nil}
	if b.head == nil {
		b.head = node
		return
	} else {
		node.next = b.head
		b.head = node
		return
	}

}

// Search a particular element in the array
func (h *HashTable) Search(k string) bool {

	index := hash(k)
	return h.array[index].search(k)

}

func (b *bucket) search(k string) bool {

	temp := b.head

	if temp == nil {
		return false
	}

	if temp.key == k {
		return true
	}

	for temp != nil {
		if temp.key == k {
			return true
		}
	}

	return false

}

// Delete elements from the Hash Table
func (h *HashTable) Delete(k string) {

	index := hash(k)
	h.array[index].delete(k)

}

func (b *bucket) delete(k string) {

	temp := b.head
	if temp == nil {
		fmt.Println("list empty")
		return
	}

	if b.head.key == k {
		b.head = b.head.next
	}

	for temp.next != nil {
		if temp.next.key == k {
			temp.next = temp.next.next
			return
		}
		temp = temp.next
	}
}

// hash function
func hash(k string) int {

	sum := 0
	for _, val := range k {
		sum += int(val)
	}
	return sum % size

}
