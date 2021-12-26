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

//XierSort 希尔排序
func XierSort(r []int) []int {
	for d := len(r) / 2; d >= 1; d = d / 2 {
		for i := d + 1; i <= len(r)-1; i++ {
			r[0] = r[i]
			j := i - d
			for j > 0 && r[0] < r[j] {
				r[j+d] = r[j]
				j = j - d
			}
			r[j+d] = r[0]
		}
	}
	return r
}

//BubbleSort 冒泡排序
func BubbleSort(r []int) []int {
	for i := 0; i < len(r)-1; i++ {
		for j := 0; j < i; j++ {
			if r[j] > r[j+1] {
				r[j], r[j+1] = r[j+1], r[j]
			}
		}
	}
	return r
}

//partition 快速排序中轴位置
func partition(r []int, first, end int) int {
	i := first
	j := end
	for i < j {
		for i < j && r[i] <= r[j] {
			j--
		}
		if i < j {
			r[i], r[j] = r[j], r[i]
			i++
		}
		for i < j && r[i] <= r[j] {
			i++
		}
		if i < j {
			r[i], r[j] = r[j], r[i]
			j--
		}
	}
	return i
}

//QuickSort 快排
func QuickSort(r []int, first, end int) []int {
	if first < end {
		pos := partition(r, first, end)
		QuickSort(r, first, pos-1)
		QuickSort(r, pos+1, end)
	}
	return r
}
