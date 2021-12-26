package main

import "fmt"

func main() {
	//iSort()
	//xSort()
	//bSort()
	//qSort()
	sSort()
}

func sSort() {
	r := []int{21, 25, 49, 28, 16, 8}
	r = SelectSort(r)
	fmt.Println(r)
}

func qSort() {
	r := []int{49, 38, 65, 97, 76, 13, 27}
	r = QuickSort(r, 0, len(r)-1)
	fmt.Println(r)
}

func bSort() {
	r := []int{5, 69, 81, 12, 38, 53, 98}
	r = BubbleSort(r)
	fmt.Println(r)
}

func iSort() {
	r := []int{0, 21, 25, 22, 10, 25, 18}
	r = InsertSort(r)
	fmt.Println(r)
}

func xSort() {
	r := []int{0, 40, 25, 49, 25, 16, 21, 8, 30, 13}
	r = XierSort(r)
	fmt.Println(r)
}
