package array

import "testing"

func Test_setZeroes(t *testing.T) {
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
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			want: args{
				matrix: [][]int{
					{1, 0, 1},
					{0, 0, 0},
					{1, 0, 1},
				},
			},
		},
		{
			name: "2",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			want: args{
				matrix: [][]int{
					{0, 0, 0, 0},
					{0, 4, 5, 0},
					{0, 3, 1, 0},
				},
			},
		},
		{
			name: "3",
			args: args{
				matrix: [][]int{
					{0, 1},
				},
			},
			want: args{
				matrix: [][]int{
					{0, 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroes(tt.args.matrix)
			if !checkRotate(tt.want.matrix, tt.args.matrix) {
				t.Errorf("setZeroes() = %v, want %v", tt.args.matrix, tt.want.matrix)
			}
		})
	}
}
