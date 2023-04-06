package main

import (
	"fmt"
	"math"
)

type Node struct {

	left *Node
	key int
	right *Node

}

type binarySearchTree struct {

	root *Node

}

func main() {

	b := &binarySearchTree{}
	b.Insert(7)
	b.Insert(3)
	b.Insert(1)
	b.Insert(8)
	b.Insert(9)
	b.Insert(43)

	// search an element in the BST
	sh := search(b.root,43)
	fmt.Println(sh)

	// checking whether a binary tree is a Binary search tree or not
	fmt.Print("In-order traversal : ")
	inOrder(b.root)
	fmt.Println()

	// delete in BST
	b.Delete(3)
	b.Delete(9)

	// checking whether a binary tree is a Binary search tree or not
	fmt.Print("In-order traversal : ")
	inOrder(b.root)
	fmt.Println()


	// Number of nodes in a binary tree
	s := getSize(b.root)
	fmt.Println("The count of the nodes: ",s)

	// Maximum element in a binary tree

	m := getMax(b.root)
	fmt.Println("The maximum value is ",m)

	// Left View of a  binary tree
	fmt.Print("The left view of the binary tree is : ")
	leftView(b.root,1)
	fmt.Println()


	// search an element using iterative in BST
	k := iSearch(b.root,43)
	fmt.Println(k)

	// check whether a given tree is a binary tree
	check := isBST(b.root,math.MinInt,math.MaxInt)
	fmt.Println("Check whether a tree is BST or not ",check)

	// find closest element in  BST
	close := closest(b.root,9)
	fmt.Println("the closest is ",close)
	

}

// Insert an element into the the tree

func (b *binarySearchTree) Insert(value int) {

		b.root = addNode(b.root,value)
}

func addNode(root *Node, value int) *Node {

	if root == nil {

		root = &Node{nil,value,nil}
		return root
	}

	if value > root.key {

		root.right = addNode(root.right,value)

	} else {

		root.left = addNode(root.left,value)

	}

	return root
}


// Delete an element in a BST

func (b *binarySearchTree) Delete(value int) {

	b.root = delete(b.root,value)
}

func delete(root *Node, value int) *Node {

	if root == nil {    
		return root      // If element is not present inside the BST
	}

	if value > root.key {

		root.right = delete(root.right,value)

	} else if value < root.key {

		root.left = delete(root.left,value)

	} else {

		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		} else {

			c := getSuc(root)
			root.key = c.key
			root.right = delete(root.right,c.key)
		}
	}

	return root
}

func getSuc(c *Node) *Node {

	c = c.right

	for c != nil && c.left != nil {
		c = c.left
	}

	return c
}

// search for an element


func search(root *Node, data int) bool {

	if root == nil {
		return false
	}

	if root.key == data {
		return true
	}

	if data > root.key {
		return search(root.right,data)
	} else {
		return search(root.left,data)
	}
}




// in order traversal
func inOrder(root *Node) {

	if root == nil {
		return
	}

	inOrder(root.left)
	fmt.Print(root.key,"  ")
	inOrder(root.right)

}


func getSize(root *Node) int {

	if root == nil {
		return 0
	}

	return 1 + getSize(root.left) + getSize(root.right)

}


// Maximum in a binary tree
func getMax(root *Node) float64 {

	if root == nil {
		return math.MinInt
	}

	return math.Max(float64(root.key),math.Max(float64(getMax(root.left)),float64(getMax(root.right))))
}


// Left View of a binary tree
var maxLevel int = 0

func leftView(root *Node,level int) {

	if root == nil {
		return
	}

	if maxLevel < level {
		fmt.Print(root.key," ")
		maxLevel = level
	}

	leftView(root.left, level + 1)
	leftView(root.right, level + 1)
}



// search for an element using iterative methods

func iSearch(root *Node, data int) bool {

	for root != nil {

		if root.key == data {
			return true
		}

		if data > root.key {
			return iSearch(root.right,data)
		} else {
			return iSearch(root.left,data)
		}
	}

	return false
}


// Check whether a Tree is BST or not

func isBST(root *Node, min int, max int) bool {

	if root == nil {
		return true
	}

	return root.key > min && root.key < max && isBST(root.left, min, root.key) && isBST(root.right,root.key, max)
}



// closest of an element
func closest(root *Node, value int) int {

	if root == nil {
		return -1
	}

	minDiff := math.MaxInt
	var currentElement int

	for root != nil {

		currentDifference := math.Abs(float64(root.key)-float64(value))

		if currentDifference < float64(minDiff) {
			minDiff = int(currentDifference)
			currentElement = root.key
		}

		if value < root.key {
			root = root.left
		} else if value > root.key {
			root = root.right
		} else {
			break
		}
	}
	return currentElement
}