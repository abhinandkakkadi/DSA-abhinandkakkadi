package main

import "fmt"

// graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

//  Vertex represents a graph vertex

type Vertex struct {
	key      int
	adjacent []*Vertex
}

// Add vertex
func (g *Graph) addVertex(k int) {

	if !contains(g.vertices, k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}

}

// Add edge
func (g *Graph) addEdge(from, to int) {

	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Println("error")
		return
	}

	fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
}

func (g *Graph) getVertex(k int) *Vertex {

	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}

	return nil
}

func contains(s []*Vertex, k int) bool {

	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {

	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v: ", v.key)
		for _, v := range v.adjacent {
			fmt.Print(" ", v.key)
		}
	}
}

func main() {

	g := Graph{}
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)
	g.addVertex(5)

	g.addEdge(1, 2)
	g.addEdge(1, 3)
	g.addEdge(1, 4)
	g.addEdge(2, 4)
	g.addEdge(2, 1)
	g.addEdge(3, 4)
	g.addEdge(4, 5)
	g.addEdge(5, 3)

	g.addEdge(2, 3)

	g.BFS(1)
	g.DFS(1)

}

// BFS TRAVERSAL
type queue []int

func (g *Graph) BFS(value int) {

	visited := make(map[int]bool)
	q := queue{}
	q = append(q, value)
	visited[value] = true

	for len(q) != 0 {

		vertex := q[0]
		q = q[1:]

		fmt.Print(" ", vertex, " ")

		for _, neighbors := range g.getVertex(vertex).adjacent {
			if !visited[neighbors.key] {
				visited[neighbors.key] = true
				q = append(q, neighbors.key)
			}
		}

	}
	fmt.Println()
}

// DFS TRAVERSAL

func (g *Graph) DFS(start int) {

	stack := []int{start}
	visited := make(map[int]bool)
	visited[start] = true

	for len(stack) > 0 {

		vertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Print(" ", vertex, " ")

		for _, neighbors := range g.getVertex(vertex).adjacent {
			if !visited[neighbors.key] {
				visited[neighbors.key] = true
				stack = append(stack, neighbors.key)
			}
		}

	}

	fmt.Println()

}
