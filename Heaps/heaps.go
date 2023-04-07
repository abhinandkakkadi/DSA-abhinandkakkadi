package main

import "fmt"

type MaxHeap struct {
	
	array []int

}

func (h *MaxHeap) Insert(key int) {

	h.array = append(h.array, key)
	h.MaxHeapify(len(h.array)-1)

}

func (h *MaxHeap) Extract(key int) int {

	extracted := h.array[0]
	l := len(h.array)-1

	if len(h.array) == 0 {
		fmt.Println("cannot extract becuase array length is 0")
		return -1
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.MaxHeapifyDown(0)

	return extracted


}


// maxHeapify will heapify from bottom top  
func (h  *MaxHeap) MaxHeapifyUp(index int) {

	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index),index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify from top to bottom

func (h *MaxHeap) MaxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l,r := leftChild(index),rightChild(index)
	childToCompare := 0

	for l <= lastIndex {

		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index,childToCompare)
			index := childToCompare
			l,r = leftChild(index),rightChild(index )
		} else {
			return
		}


	}	

}

// get the parent index
func parent(i int) int {
	return (i-1)/2
}

// get the left child
func leftChild(i int) int {
	return 2*i + 1
}

// get the right child
func rightChild(i int) int {
	return 2*i + 2
}

// swap two nodes
func (h *MaxHeap) swap(i1,i2 int) {

	h.array[i1],h.array[i2] = h.array[i2],h.array[i1]

}

func main() {

	h := &MaxHeap{}
	h.Insert(23)
	h.Insert(34)
	h.Insert(22)

	fmt.Println(h)
}