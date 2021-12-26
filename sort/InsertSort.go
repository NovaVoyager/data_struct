package main

//InsertSort 直接插入排序
func InsertSort(r []int) []int {
	for i := 2; i <= len(r)-1; i++ {
		r[0] = r[i]
		j := i - 1
		for r[0] < r[j] {
			r[j+1] = r[j]
			j--
		}
		r[j+1] = r[0]
	}
	return r
}
