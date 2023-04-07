package main

import "fmt"

// MAX HEAP

type MaxHeap struct {
	
	array []int

}


func (h *MaxHeap) Insert(key int) {

	h.array = append(h.array, key)
	h.MaxHeapifyUp(len(h.array)-1)

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


// MIN HEAP

type MinHeap struct {
	array []int
}

// Insert value in Min heap
func (mH *MinHeap) Insert(value int) {

	mH.array = append(mH.array, value)
	mH.MinHeapifyUp(len(mH.array)-1)

}

func (mH *MinHeap) MinHeapifyUp(index int) {

	for mH.array[index] < mH.array[parent(index)] {
		mH.swap(index,parent(index))
		index = parent(index)
	}
}

//	Extract value int Min Heap

func (mH *MinHeap) Extract() {

	l := len(mH.array) - 1
	mH.array[0] = mH.array[l]
	mH.array = mH.array[:l]
	mH.MinHeapifyDown(0)
}

func (mH *MinHeap) MinHeapifyDown(index int) {

	lastIndex := len(mH.array) - 1
	l,r := leftChild(index),rightChild(index)
	compareChild := 0

	for l <= lastIndex {

		if l == lastIndex {
			compareChild = l
		} else if mH.array[l] > mH.array[r] {
			compareChild = l
		} else {
			compareChild = r
		}

		if mH.array[index] > mH.array[compareChild] {
			mH.swap(index,compareChild)
			index = compareChild
			l,r = leftChild(index),rightChild(index)
		} else {
			return
		}
	}
}



// COMMON OPERATIONS

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

// swap two nodes MaxHeap
func (h *MaxHeap) swap(i1,i2 int) {

	h.array[i1],h.array[i2] = h.array[i2],h.array[i1]

}

// swap two nodes MinHeap
func (mH *MinHeap) swap(i1, i2 int) {

	mH.array[i1],mH.array[i2] = mH.array[i2],mH.array[i1]

}

func main() {

	h := &MaxHeap{}
	h.Insert(23)
	h.Insert(34)
	h.Insert(22)
	h.Insert(2)

	fmt.Println(h)

	minH := &MinHeap{}
	minH.Insert(23)
	minH.Insert(34)
	minH.Insert(22)
	minH.Insert(2)
	fmt.Println(minH)

	minH.Extract()
	fmt.Println(minH)

}