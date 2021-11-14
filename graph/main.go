package main

import "fmt"

func main() {
	elems := []Elem{
		{V1: "A", V2: "B", Weight: 10},
		{V1: "A", V2: "C", Weight: 12},
		{V1: "A", V2: "E", Weight: 5},
		{V1: "B", V2: "D", Weight: 20},
		{V1: "C", V2: "D", Weight: 25},
	}
	vexs := []VerTextType{"A", "B", "C", "D", "E"}
	graph := NewAMGraph(vexs, elems)
	fmt.Println(graph)
}
