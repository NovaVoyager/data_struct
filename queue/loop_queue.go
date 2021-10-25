package main

type MyLoopCircularQueue struct {
	data       []int
	head, tail int
	size       int
}

func ConstructorLoop(k int) MyLoopCircularQueue {
	return MyLoopCircularQueue{
		data: make([]int, k, k),
		head: -1,
		tail: -1,
		size: k,
	}
}

//EnQueue 入队
func (this *MyLoopCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	if this.IsEmpty() {
		this.head = 0 //入队第一个元素时，把开始元素下标为0
	}
	this.tail = (this.tail + 1) % this.size //计算出下一个下标位置
	this.data[this.tail] = value
	return true
}

//DeQueue 出队
func (this *MyLoopCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.head == this.tail {
		this.head = -1
		this.tail = -1
		return true
	}
	this.head = (this.head + 1) % this.size
	return true
}

//Front 获取对首元素
func (this *MyLoopCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

//Rear 获取队尾元素
func (this *MyLoopCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.tail]
}

//IsEmpty 队列是否为空
func (this *MyLoopCircularQueue) IsEmpty() bool {
	return this.head == -1
}

//IsFull 队列是否满
func (this *MyLoopCircularQueue) IsFull() bool {
	return ((this.tail + 1) % this.size) == this.head
}
