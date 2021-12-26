package main

import "fmt"

func main() {
	iSort()
}

func iSort() {
	r := []int{0, 21, 25, 22, 10, 25, 18}
	r = InsertSort(r)
	fmt.Println(r)
}
