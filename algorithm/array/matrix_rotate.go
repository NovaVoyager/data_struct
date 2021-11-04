package array

//rotate 旋转矩阵
func rotate1(matrix [][]int) {
	x := len(matrix) - 1
	l := len(matrix)

	newMatrix := make([][]int, 0, 3)
	for i := 0; i < l; i++ {
		in := make([]int, l, l)
		newMatrix = append(newMatrix, in)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			newMatrix[j][x] = matrix[i][j]
		}
		x--
	}
	copy(matrix, newMatrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n-1-i] = v
		}
	}
	copy(matrix, tmp) // 拷贝 tmp 矩阵每行的引用
}
