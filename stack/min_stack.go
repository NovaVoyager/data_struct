package main

type MinStack struct {
	top *Node
}

type Node struct {
	val, min int
	next     *Node
}

func Constructor() MinStack {
	return MinStack{}
}

//Push 入栈
func (this *MinStack) Push(val int) {
	node := Node{val: val}

	if this.top == nil {
		node.min = val
	} else {
		node.min = min(val, this.top.min)
	}

	node.next = this.top
	this.top = &node
}

func (this *MinStack) Pop() {
	if this.IsEmpty() {
		return
	}
	this.top = this.top.next
}

func (this *MinStack) Top() int {
	return this.top.val
}

func (this *MinStack) GetMin() int {
	return this.top.min
}

func (this *MinStack) IsEmpty() bool {
	if this.top == nil {
		return true
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
