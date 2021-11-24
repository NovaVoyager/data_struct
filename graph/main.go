package main

import "fmt"

func main() {
	dijkstra()
}

func dijkstra() {
	vexs := []VerTextType{"A", "B", "C", "D", "E", "F", "G"}
	elems := []Elem{
		{V1: "A", V2: "B", Weight: 4},
		{V1: "A", V2: "C", Weight: 6},
		{V1: "A", V2: "D", Weight: 6},
		{V1: "B", V2: "C", Weight: 1},
		{V1: "B", V2: "E", Weight: 7},
		{V1: "C", V2: "E", Weight: 6},
		{V1: "C", V2: "F", Weight: 4},
		{V1: "D", V2: "C", Weight: 2},
		{V1: "D", V2: "F", Weight: 5},
		{V1: "E", V2: "G", Weight: 6},
		{V1: "F", V2: "E", Weight: 1},
		{V1: "F", V2: "G", Weight: 8},
	}
	graph := NewAMGraph(vexs, elems)
	fmt.Println(graph)
}

func edge() {
	vertexs := []VerTextType{"A", "B", "C", "D", "E", "F"}
	edges := []EdgeType{
		{from: 0, to: 1, weight: 30},
		{from: 0, to: 2, weight: 46},
		{from: 0, to: 5, weight: 19},
		{from: 1, to: 4, weight: 12},
		{from: 2, to: 5, weight: 25},
		{from: 2, to: 3, weight: 17},
		{from: 3, to: 4, weight: 38},
		{from: 4, to: 5, weight: 26},
		{from: 3, to: 5, weight: 25},
	}

	g := NewEdgeGraph(vertexs, edges)
	g.Kruskal()
}

func am() {
	vexs := []VerTextType{"A", "B", "C", "D", "E"}
	elems := []Elem{
		{V1: "A", V2: "B", Weight: 10},
		{V1: "A", V2: "C", Weight: 12},
		{V1: "A", V2: "E", Weight: 5},
		{V1: "B", V2: "D", Weight: 20},
		{V1: "B", V2: "E", Weight: 18},
		{V1: "C", V2: "D", Weight: 25},
	}
	graph := NewNoAMGraph(vexs, elems)
	//graph.forDepthAMGraph(1)
	graph.Prim(0)
}

func al() {
	elems := []ALGraphElem{
		{
			vex: "A",
			side: []ALElem{
				{V1: "B", Weight: 10},
				{V1: "D", Weight: 30},
			},
		},
		{
			vex: "B",
			side: []ALElem{
				{V1: "E", Weight: 20},
				{V1: "C", Weight: 1},
			},
		},
		{
			vex: "C",
			side: []ALElem{
				{V1: "D", Weight: 2},
				{V1: "E", Weight: 3},
			},
		},
	}
	vexs := []VerTextType{"A", "B", "C", "D", "E"}
	g := NewALGraph(vexs, elems)
	g.forBreadthALGraph()
}
