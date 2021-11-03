package array

import "sort"

//merge 合并区间
//eg:1
//输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出：[[1,6],[8,10],[15,18]]
//解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
//eg:2
//输入：intervals = [[1,4],[4,5]]
//输出：[[1,5]]
//解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	tmp := intervals[0]
	for _, v := range intervals {
		if v[0] <= tmp[1] && v[0] >= tmp[0] {
			left := min(v[0], tmp[0])
			right := max(v[1], tmp[1])
			tmp = []int{left, right}
		} else {
			result = append(result, tmp)
			tmp = v
		}
	}
	result = append(result, tmp)
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
