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
			return node.Val
		}
		node = node.Next
	}
	return -1
}

//AddAtHead 头插
func (this *MyLinkedList) AddAtHead(val int) {
	//第一个节点
	if this.Size == 0 {
		this.pushFirstVal(val)
		return
	}

	newNode := DoubleLinked{
		Val:  val,
		Prev: this.Head.Prev,
		Next: this.Head,
	}
	this.Head.Prev = &newNode
	this.Head = &newNode
	this.Size++
}

//AddAtTail 尾插
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Size == 0 {
		this.pushFirstVal(val)
		return
	}
	node := this.Head
	for node != nil {
		if node.Next != nil {
			node = node.Next
			continue
		}
		newNode := DoubleLinked{
			Val:  val,
			Prev: node,
			Next: node.Next,
		}
		node.Next = &newNode
		this.Size++
	}
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
		}
		newNode := DoubleLinked{
			Val:  val,
			Prev: node.Prev,
			Next: node,
		}
		if node.Prev != nil {
			node.Prev.Next = &newNode
		}
		node.Prev = &newNode
		this.Size++
		node = nil
	}
}

//DeleteAtIndex 删除指定位置节点
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if !this.isValidIndex(index) {
		return
	}

	//删除第一个节点
	if index == 0 {
		if this.Head.Prev != nil {
			this.Head.Next.Prev = this.Head.Prev
		}
		this.Head = this.Head.Next
		this.Size--
	}

	//删除指定位置节点
	node := this.Head
	for i := 0; node != nil; i++ {
		if i == index && i == this.Size-1 { //删除最后一个节点
			node.Prev.Next = node.Next
			node.Prev = nil
			this.Size--
			return
		} else if i == index { //删除指定位置
			node.Prev.Next = node.Next
			node.Prev, node.Next = nil, nil
			this.Size--
			return
		}
		node = node.Next
	}
}

//isValidIndex 索引是否有效  true 有效，false 无效
func (this *MyLinkedList) isValidIndex(index int) bool {
	if index < 0 || index > this.Size-1 {
		return false
	}
	return true
}

//pushFirstVal 插入第一个元素
func (this *MyLinkedList) pushFirstVal(val int) {
	newNode := DoubleLinked{
		Val:  val,
		Prev: nil,
		Next: nil,
	}
	this.Head = &newNode
	this.Size++
}
