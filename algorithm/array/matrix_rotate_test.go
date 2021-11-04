package array

import "testing"

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{
			name: "1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: args{
				matrix: [][]int{
					{7, 4, 1},
					{8, 5, 2},
					{9, 6, 3},
				},
			},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{
					{5, 1, 9, 11},
					{2, 4, 8, 10},
					{13, 3, 6, 7},
					{15, 14, 12, 16},
				},
			},
			want: args{
				matrix: [][]int{
					{15, 13, 2, 5},
					{14, 3, 4, 1},
					{12, 6, 8, 9},
					{16, 7, 10, 11},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate1(tt.args.matrix)
			if !checkRotate(tt.want.matrix, tt.args.matrix) {
				t.Errorf("rotate() = %v, want %v", tt.args.matrix, tt.want.matrix)
			}
		})
	}
}

func checkRotate(x, y [][]int) bool {
	l := len(x)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if x[i][j] != y[i][j] {
				return false
			}
		}
	}
	return true
}
