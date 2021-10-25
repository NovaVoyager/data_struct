package main

type MyCircularQueue struct {
	head, tail   *Node
	size, queLen int
}

type Node struct {
	data int
	next *Node
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		head:   nil,
		tail:   nil,
		size:   k,
		queLen: 0,
	}
}

//EnQueue 入队
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	node := Node{data: value}

	if this.IsEmpty() {
		this.head = &node
		this.tail = &node
		this.queLen++
		return true
	}

	this.tail.next = &node
	this.tail = &node
	this.queLen++
	return true
}

//DeQueue 出队
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	if this.head.next == nil {
		this.head, this.tail = nil, nil
		this.queLen = 0
		return true
	}

	this.head = this.head.next
	this.queLen--
	return true
}

//Front 获取对首元素
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.data
}

//Rear 获取队尾元素
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.data
}

//IsEmpty 队列是否为空
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == nil && this.tail == nil
}

//IsFull 队列是否满
func (this *MyCircularQueue) IsFull() bool {
	return this.queLen == this.size
}
