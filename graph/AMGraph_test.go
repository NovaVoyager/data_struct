package main

import (
	"reflect"
	"testing"
)

func TestNewAMGraph(t *testing.T) {
	type args struct {
		vexs  []VerTextType
		elems []Elem
	}
	tests := []struct {
		name string
		args args
		want *AMGraph
	}{
		{
			name: "1",
			args: args{
				vexs: []VerTextType{"A", "B", "C", "D", "E"},
				elems: []Elem{
					{V1: "A", V2: "B", Weight: 10},
					{V1: "A", V2: "C", Weight: 12},
					{V1: "A", V2: "E", Weight: 5},
					{V1: "B", V2: "D", Weight: 20},
					{V1: "C", V2: "D", Weight: 25},
				},
			},
			want: &AMGraph{
				vexs: []VerTextType{"A", "B", "C", "D", "E"},
				arcs: [][]ArcType{
					{32767, 10, 12, 32767, 5},
					{10, 32767, 32767, 20, 32767},
					{12, 32767, 32767, 25, 32767},
					{32767, 20, 25, 32767, 32767},
					{5, 32767, 32767, 32767, 32767},
				},
				vexNum: 5,
				arcNum: 5,
			},
		},
		{
			name: "2",
			args: args{
				vexs: []VerTextType{"A", "B", "C", "D", "E", "F"},
				elems: []Elem{
					{V1: "A", V2: "B", Weight: 5},
					{V1: "A", V2: "F", Weight: 3},
					{V1: "A", V2: "D", Weight: 7},
					{V1: "A", V2: "C", Weight: 8},

					{V1: "B", V2: "C", Weight: 4},
					{V1: "C", V2: "F", Weight: 9},
					{V1: "C", V2: "D", Weight: 5},

					{V1: "D", V2: "E", Weight: 5},
					{V1: "D", V2: "F", Weight: 6},

					{V1: "E", V2: "F", Weight: 1},
				},
			},
			want: &AMGraph{
				vexs: []VerTextType{"A", "B", "C", "D", "E", "F"},
				arcs: [][]ArcType{
					{32767, 5, 8, 7, 32767, 3},
					{5, 32767, 4, 32767, 32767, 32767},
					{8, 4, 32767, 5, 32767, 9},
					{7, 32767, 5, 32767, 5, 6},
					{32767, 32767, 32767, 5, 32767, 1},
					{3, 32767, 9, 6, 1, 32767},
				},
				vexNum: 6,
				arcNum: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoAMGraph(tt.args.vexs, tt.args.elems); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAMGraph() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}

func TestAMGraph_minEdge(t *testing.T) {
	type fields struct {
		vexs   []VerTextType
		arcs   [][]ArcType
		vexNum int
		arcNum int
	}
	type args struct {
		se []shortEdge
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "1",
			fields: fields{
				vexs: []VerTextType{"A", "B", "C", "D", "E", "F"},
				arcs: [][]ArcType{
					{0, 34, 46, MaxInt, MaxInt, 19},
					{34, 0, MaxInt, MaxInt, 12, MaxInt},
					{46, MaxInt, 0, 17, MaxInt, 25},
					{MaxInt, MaxInt, 17, 0, 38, 25},
					{MaxInt, 12, MaxInt, 38, 0, 26},
					{19, MaxInt, 25, 25, 26, 0},
				},
				vexNum: 6,
				arcNum: 6,
			},
			args: args{
				se: []shortEdge{
					{adjvex: 0, lowCost: 0},
					{adjvex: 0, lowCost: 34},
					{adjvex: 0, lowCost: 46},
					{adjvex: 0, lowCost: MaxInt},
					{adjvex: 0, lowCost: MaxInt},
					{adjvex: 0, lowCost: 19},
				},
			},
			want: 5,
		},
		{
			name: "2",
			fields: fields{
				vexs: []VerTextType{"A", "B", "C", "D", "E", "F"},
				arcs: [][]ArcType{
					{0, 34, 46, MaxInt, MaxInt, 19},
					{34, 0, MaxInt, MaxInt, 12, MaxInt},
					{46, MaxInt, 0, 17, MaxInt, 25},
					{MaxInt, MaxInt, 17, 0, 38, 25},
					{MaxInt, 12, MaxInt, 38, 0, 26},
					{19, MaxInt, 25, 25, 26, 0},
				},
				vexNum: 6,
				arcNum: 6,
			},
			args: args{
				se: []shortEdge{
					{adjvex: 0, lowCost: 0},
					{adjvex: 0, lowCost: 34},
					{adjvex: 5, lowCost: 25},
					{adjvex: 5, lowCost: 25},
					{adjvex: 5, lowCost: 26},
					{adjvex: 0, lowCost: 0},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &AMGraph{
				vexs:   tt.fields.vexs,
				arcs:   tt.fields.arcs,
				vexNum: tt.fields.vexNum,
				arcNum: tt.fields.arcNum,
			}
			if got := this.minEdge(tt.args.se); got != tt.want {
				t.Errorf("minEdge() = %v, want %v", got, tt.want)
			}
		})
	}
}
