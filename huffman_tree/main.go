package main

func main() {
	n := 8
	//ht := createTree(n, []int{5, 29, 7, 8, 14, 23, 3, 11})
	//ht := createTree(n, []int{40, 30, 15, 5, 4, 3, 3})
	ht := createTree(n, []int{5, 29, 7, 8, 14, 23, 3, 11})
	//_ = huffmanCode(ht, n, []ElemType{"", "A", "B", "C", "D", "E", "F", "G", "H"})
	_ = huffmanCode(ht, n, []ElemType{"", "A", "B", "C", "D", "E", "F", "G", "H"})
}
