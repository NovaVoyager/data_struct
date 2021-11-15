package main

import "fmt"

func main() {
	am()
}

func am() {
	vexs := []VerTextType{"A", "B", "C", "D", "E"}
	elems := []Elem{
		{V1: "A", V2: "B", Weight: 10},
		{V1: "A", V2: "C", Weight: 12},
		{V1: "A", V2: "E", Weight: 5},
		{V1: "B", V2: "D", Weight: 20},
		{V1: "C", V2: "D", Weight: 25},
	}
	graph := NewNoAMGraph(vexs, elems)
	graph.forDepthAMGraph(1)
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
	fmt.Println(g)
}
