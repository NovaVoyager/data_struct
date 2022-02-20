package linkList

//MyLinkedListOne linked defined
type MyLinkedListOne struct {
	Head *ListNode
	Size int
}

//ListNode defined
type ListNode struct {
	val  int
	Next *ListNode
}

func ConstructorOne() MyLinkedListOne {
	return MyLinkedListOne{}
}

//Get 根据索引获取值
func (this *MyLinkedListOne) Get(index int) int {
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
func (this *MyLinkedListOne) AddAtHead(val int) {
	node := ListNode{
		val:  val,
		Next: this.Head,
	}
	this.Size++
	this.Head = &node
}

//AddAtTail 尾插
func (this *MyLinkedListOne) AddAtTail(val int) {
	newNode := ListNode{
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
func (this *MyLinkedListOne) AddAtIndex(index int, val int) {
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
			newNode := ListNode{
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
func (this *MyLinkedListOne) DeleteAtIndex(index int) {
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
func (this *MyLinkedListOne) IsValidIndex(index int) bool {
	if index < 0 || index > this.Size-1 {
		return false
	}
	return true
}
