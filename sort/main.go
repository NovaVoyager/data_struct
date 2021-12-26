package main

import "fmt"

func main() {
	//iSort()
	xSort()
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
