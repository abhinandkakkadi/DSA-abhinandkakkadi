package main

import "fmt"

// graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

//  Vertex represents a graph vertex

type Vertex struct {
	key int
	adjacent []*Vertex
}

// Add vertex
func (g *Graph) AddVertex(k int) {

	if !contains(g.vertices,k) {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
	
}

// Add edge
func (g *Graph) AddEdge(from,to int) {

	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Println("error")
		return
	}

	fromVertex.adjacent = append(fromVertex.adjacent,toVertex)
}

func (g *Graph) getVertex(k int) *Vertex {

	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}

	return nil
}



func contains(s []*Vertex,k int) bool {

	for _,v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}


func (g *Graph) Print() {

	for _,v := range g.vertices {
		fmt.Printf("\nVertex %v: ",v.key)
		for _,v := range v.adjacent {
			fmt.Print(" ",v.key)
		}
	}
}

func main() {

	test := &Graph{}


	for i:=0; i<5; i++ {
		test.AddVertex(i)
	}

	test.AddEdge(0,1)
	test.AddEdge(0,2)
	test.AddEdge(1,2)
	test.AddEdge(4,3)
	test.Print()

}
