package array

func setZeroes(matrix [][]int) {
	rows := make([]int, 0)
	cols := make([]int, 0)
	for i, row := range matrix {
		for j, v := range row {
			if v == 0 {
				if !inArray(i, rows) {
					rows = append(rows, i)
				}
				if !inArray(j, cols) {
					cols = append(cols, j)
				}
			}
		}
	}

	n := len(matrix)
	for _, row := range rows {
		for j, _ := range matrix[row] {
			matrix[row][j] = 0
		}
	}

	for _, v := range cols {
		for j := 0; j < n; j++ {
			matrix[n-j-1][v] = 0
		}
	}
}

func inArray(x int, arr []int) bool {
	tmp := make(map[int]struct{}, len(arr))
	for _, v := range arr {
		tmp[v] = struct{}{}
	}
	_, ok := tmp[x]
	return ok
}
