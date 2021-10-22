package main

//MyLinkedList 双链表
type MyLinkedList struct {
	Head *DoubleLinked
	Size int
}

//DoubleLinked 双链表节点
type DoubleLinked struct {
	Val        int
	Prev, Next *DoubleLinked
}

//Constructor 初始化双向链表
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

//Get 获取指定索引的值
func (this *MyLinkedList) Get(index int) int {
	if !this.isValidIndex(index) {
		return -1
	}
	node := this.Head
	for i := 0; node != nil; i++ {
		if i == index {
			break
		}
		node = node.Next
	}
	return node.Val
}

//AddAtHead 头插
func (this *MyLinkedList) AddAtHead(val int) {
	node := DoubleLinked{Val: val}
	if this.Size == 0 { //第一个节点
		this.Head = &node
	} else {
		node.Next = this.Head
		this.Head.Prev = &node
		this.Head = &node
	}
	this.Size++
}

//AddAtTail 尾插
func (this *MyLinkedList) AddAtTail(val int) {
	node := DoubleLinked{Val: val}
	if this.Size == 0 { //第一个节点
		this.Head = &node
	} else {
		cur := this.Head
		for cur != nil {
			if cur.Next != nil {
				cur = cur.Next
				continue
			} else {
				cur.Next = &node
				node.Prev = cur
				break
			}
		}
	}
	this.Size++
}

//AddAtIndex 指定位置插入节点
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	if index > this.Size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	node := this.Head
	for i := 0; node != nil; i++ {
		if i != index {
			node = node.Next
			continue
		} else {
			newNode := DoubleLinked{
				Val:  val,
				Prev: node.Prev,
				Next: node,
			}
			node.Prev.Next = &newNode
			node.Prev = &newNode
			this.Size++
			break
		}
	}
}

//DeleteAtIndex 删除指定位置节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if !this.isValidIndex(index) {
		return
	}

	//删除第一个节点
	if index == 0 {
		this.Head = this.Head.Next
		this.Size--
		return
	}

	//删除指定位置节点
	node := this.Head
	for i := 0; node != nil; i++ {
		if index == 0 && this.Size == 1 { //把仅有的一个节点删除
			this.Head = nil
			this.Size = 0
		} else if i == index && i == this.Size-1 { //删除链表最尾部的一个节点
			node.Prev.Next = node.Next
			node.Prev = nil
			this.Size--
		} else if i == index { //删除指定位置
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
			node.Prev, node.Next = nil, nil
			this.Size--
		}
		node = node.Next
	}
}

//isValidIndex 索引是否有效  true 有效，false 无效
func (this *MyLinkedList) isValidIndex(index int) bool {
	if index < 0 || index > this.Size-1 || this.Size == 0 {
		return false
	}
	return true
}
