package array

//findDiagonalOrder 对角线遍历
func findDiagonalOrder(mat [][]int) []int {
	length := len(mat)
	count := length + length - 1 //循环次数
	x, y := 0, 0
	result := make([]int, 0, length*length)

	//i%2
	//向上为奇数
	//向下为偶数
	for i := 0; i < count; i++ {
		if i%2 == 0 { //向上
			for x >= 0 && y < length {
				result = append(result, mat[x][y])
				x--
				y++
			}
			if y < length {
				x++
			} else {
				x = x + 2
				y--
			}
		} else { //向下
			for x < length && y >= 0 {
				result = append(result, mat[x][y])
				x++
				y++
			}
			if x < length {
				y++
			} else {
				x--
				y = y + 2
			}
		}
	}

	return result
}
