package main

//MyLinkedList linked defined
type MyLinkedList struct {
	Head *Linked
	Size int
}

//Linked defined
type Linked struct {
	val  int
	Next *Linked
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	return 0
}

func (this *MyLinkedList) AddAtHead(val int) {

}

func (this *MyLinkedList) AddAtTail(val int) {

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

}

func main() {

}
