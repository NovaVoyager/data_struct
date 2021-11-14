package main

import "fmt"

func main() {
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
