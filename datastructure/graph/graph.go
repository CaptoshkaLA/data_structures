package graph

import (
	"fmt"
)

type Graph struct {
	vertices map[interface{}]map[interface{}]struct{}
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[interface{}]map[interface{}]struct{}),
	}
}

func (g *Graph) AddVertex(vertex interface{}) {
	if _, exists := g.vertices[vertex]; !exists {
		g.vertices[vertex] = make(map[interface{}]struct{})
	}
}

func (g *Graph) AddEdge(vertex1, vertex2 interface{}) {
	g.AddVertex(vertex1)
	g.AddVertex(vertex2)
	g.vertices[vertex1][vertex2] = struct{}{}
	g.vertices[vertex2][vertex1] = struct{}{}
}

func (g *Graph) PrintGraph() {
	fmt.Println("Список вершин и связей в графе:")
	for vertex, edges := range g.vertices {
		fmt.Printf("%v:", vertex)
		for edge := range edges {
			fmt.Printf(" %v", edge)
		}
		fmt.Println()
	}
}

func TestCaseGraph() {
	fmt.Println("\nGraph\n")

	// Создаем новый граф
	graph := NewGraph()

	// Добавляем вершины
	graph.AddVertex("A")
	graph.AddVertex("B")
	graph.AddVertex("C")
	graph.AddVertex("D")

	// Добавляем рёбра
	graph.AddEdge("A", "B")
	graph.AddEdge("B", "C")
	graph.AddEdge("C", "D")
	graph.AddEdge("D", "A")

	// Печатаем граф
	graph.PrintGraph()
}
