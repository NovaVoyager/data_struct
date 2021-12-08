package main

import (
	"fmt"
	"github.com/miaogu-go/data_struct/graph/CriticalPathAlgorithm"
)

func main() {
	alV()
}

func floyd() {
	vexs := []VerTextType{"A", "B", "C", "D"}
	elems := []Elem{
		{V1: "A", V2: "B", Weight: 1},
		{V1: "A", V2: "D", Weight: 4},
		{V1: "B", V2: "C", Weight: 9},
		{V1: "B", V2: "D", Weight: 2},
		{V1: "C", V2: "A", Weight: 3},
		{V1: "C", V2: "B", Weight: 5},
		{V1: "C", V2: "D", Weight: 8},
		{V1: "D", V2: "C", Weight: 6},
	}
	graph := NewAMGraph(vexs, elems)
	graph.Floyd()
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
	graph.Dijkstra(0)
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

func alt() {
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
	vexs := []Vexs{
		{V: "A"},
		{V: "B"},
		{V: "C"},
		{V: "D"},
		{V: "E"},
	}
	g := NewALGraph(vexs, elems)
	g.forBreadthALGraph()
}

func alAOV() {
	elems := []CriticalPathAlgorithm.ALGraphElem{
		{
			Vex:  "A",
			Side: nil,
		},
		{
			Vex: "B",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "A"},
				{V1: "D"},
			},
		},
		{
			Vex: "C",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "A"},
				{V1: "D"},
			},
		},
		{
			Vex: "D",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "A"},
				{V1: "F"},
			},
		},
		{
			Vex: "E",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "C"},
				{V1: "D"},
				{V1: "F"},
			},
		},
		{
			Vex:  "F",
			Side: nil,
		},
	}
	vexs := []CriticalPathAlgorithm.Vexs{
		{V: "A", InNum: 3},
		{V: "B", InNum: 0},
		{V: "C", InNum: 1},
		{V: "D", InNum: 3},
		{V: "E", InNum: 0},
		{V: "F", InNum: 2},
	}
	g := CriticalPathAlgorithm.NewALGraph(vexs, elems)
	g.AOVSort()
}

func alV() {
	elems := []CriticalPathAlgorithm.ALGraphElem{
		{
			Vex: "V0",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "V1", Weight: 6},
				{V1: "V2", Weight: 4},
				{V1: "V3", Weight: 5},
			},
		},
		{
			Vex: "V1",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "V4", Weight: 1},
			},
		},
		{
			Vex: "V2",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "V4", Weight: 1},
			},
		},
		{
			Vex: "V3",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "V5", Weight: 2},
			},
		},
		{
			Vex: "V4",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "V6", Weight: 9},
				{V1: "V7", Weight: 7},
			},
		},
		{
			Vex: "V5",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "V7", Weight: 4},
			},
		},
		{
			Vex: "V6",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "V8", Weight: 2},
			},
		},
		{
			Vex: "V7",
			Side: []CriticalPathAlgorithm.ALElem{
				{V1: "V8", Weight: 4},
			},
		},
		{
			Vex:  "V8",
			Side: nil,
		},
	}
	vexs := []CriticalPathAlgorithm.Vexs{
		{V: "V0", InNum: 0},
		{V: "V1", InNum: 1},
		{V: "V2", InNum: 1},
		{V: "V3", InNum: 1},
		{V: "V4", InNum: 2},
		{V: "V5", InNum: 1},
		{V: "V6", InNum: 1},
		{V: "V7", InNum: 2},
		{V: "V8", InNum: 2},
	}
	g := CriticalPathAlgorithm.NewALGraph(vexs, elems)
	g.AOVSort()
	err := g.CriticalPath()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
