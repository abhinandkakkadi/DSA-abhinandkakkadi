package main

import "fmt"

// Node represents each node in the trie
const size = 26

type Node struct {
	children [size]*Node
	isEnd bool
}


// Trie
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert
func (t *Trie) Insert(w string) {

	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}


// Search



func main() {

	testTrie := InitTrie()
	fmt.Println(testTrie.root)

}