package main

import "fmt"

// Node represents each node in the trie
const size = 26

type Node struct {
	children [size]*Node
	isEnd    bool
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
func (t *Trie) Search(w string) bool {

	currentNode := t.root

	for i := 0; i < len(w); i++ {

		currentIndex := w[i] - 'a'
		if currentNode.children[currentIndex] == nil {
			return false
		}
		currentNode = currentNode.children[currentIndex]

	}

	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {

	t := InitTrie()
	t.Insert("abhinand")
	t.Insert("athira")

	fmt.Println(t.Search("abhinand"))
	fmt.Println(t.Search("athira"))

}
