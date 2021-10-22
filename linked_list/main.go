package main

import "fmt"

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

//Get 根据索引获取值
func (this *MyLinkedList) Get(index int) int {
	if !this.IsValidIndex(index) { //索引为负数或大于链表最大索引，则为无效
		return -1
	}
	node := this.Head
	for i := 0; node != nil; i++ {
		if index == i {
			return node.val
		}
		node = node.Next
	}
	return -1
}

//AddAtHead 头插
func (this *MyLinkedList) AddAtHead(val int) {
	node := Linked{
		val:  val,
		Next: this.Head,
	}
	this.Size++
	this.Head = &node
}

//AddAtTail 尾插
func (this *MyLinkedList) AddAtTail(val int) {
	newNode := Linked{
		val:  val,
		Next: nil,
	}

	if this.Size == 0 { //链表为空
		this.Head = &newNode
		this.Size++
		return
	}

	node := this.Head
	for node != nil {
		if node.Next != nil {
			node = node.Next
			continue
		}
		node.Next = &newNode
		this.Size++
		node = nil
	}
}

//AddAtIndex 指定索引位置前插入节点
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size { //大于链表长度，则不会插入节点
		return
	}
	if index < 0 { //小于0，则在头部插入节点
		this.AddAtHead(val)
		return
	}
	if index == this.Size { //等于链表的长度，则该节点将附加到链表的末尾
		this.AddAtTail(val)
		return
	}

	index--        //插入在某个索引前，也就是插入在这个索引-1的后面
	if index < 0 { //插入在第一个元素的前面，使用头插
		this.AddAtHead(val)
		return
	}

	node := this.Head
	for i := 0; node != nil; i++ {
		if i == index {
			newNode := Linked{
				val:  val,
				Next: node.Next,
			}
			node.Next = &newNode
			this.Size++
			return
		}
		node = node.Next
	}
}

//DeleteAtIndex 删除指定位置索引节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if !this.IsValidIndex(index) {
		return
	}
	if index == 0 { //删除第一个节点
		this.Head = this.Head.Next
		this.Size--
		return
	}

	index--
	node := this.Head
	for i := 0; node != nil; i++ {
		if i == this.Size-2 && index == this.Size-2 { //删除最后一个节点
			node.Next = nil
			this.Size--
			return
		} else if index == i { //删除指定位置节点
			node.Next = node.Next.Next
			this.Size--
			return
		}
		node = node.Next
	}
}

//IsValidIndex 索引是否有效 true 有效，false 无效
func (this *MyLinkedList) IsValidIndex(index int) bool {
	if index < 0 || index > this.Size-1 {
		return false
	}
	return true
}

func main() {
	linked := Constructor()
	linked.AddAtIndex(0, 10)
	linked.AddAtIndex(0, 20)
	linked.AddAtIndex(1, 30)
	linked.AddAtIndex(1, 40)
	linked.AddAtIndex(1, 50)
	linked.DeleteAtIndex(2)
	linked.DeleteAtIndex(0)
	linked.DeleteAtIndex(2)
	/*linked.AddAtHead(1)
	linked.AddAtHead(2)
	linked.AddAtTail(3)
	linked.AddAtIndex(2, 4)
	linked.AddAtIndex(-1, 0)
	linked.AddAtIndex(5, 5)
	linked.AddAtIndex(7, 7)
	linked.DeleteAtIndex(5)
	linked.DeleteAtIndex(0)
	linked.DeleteAtIndex(2)
	linked.DeleteAtIndex(1)
	linked.DeleteAtIndex(-1)
	linked.DeleteAtIndex(2)*/
	fmt.Println(linked)
}
