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

func rotate2(matrix [][]int) {
	n := len(matrix)
	newMatrix := make([][]int, 0, n)
	for _, _ = range matrix {
		newMatrix = append(newMatrix, make([]int, n, n))
	}

	for i, row := range matrix {
		for j, v := range row {
			newMatrix[j][n-i-1] = v
		}
	}

	copy(matrix, newMatrix)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

func rotate3(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
