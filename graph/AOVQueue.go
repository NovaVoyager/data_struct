package main

type AOVQueueNode struct {
	data VexNode
	next *AOVQueueNode
}

type AOVQueue struct {
	first *AOVQueueNode
	last  *AOVQueueNode
}

//NewAOVQueue 初始化队列
func NewAOVQueue() *AOVQueue {
	return &AOVQueue{}
}

//pushQueue 入队
func (this *AOVQueue) pushQueue(data VexNode) {
	tmp := &AOVQueueNode{
		data: data,
		next: nil,
	}

	if this.first == nil && this.last == nil {
		this.first = tmp
		this.last = tmp
		return
	}

	this.last.next = tmp
	this.last = tmp
}

//popQueue 出队
func (this *AOVQueue) popQueue() VexNode {
	tmp := this.first
	this.first = tmp.next
	if tmp.next == nil {
		this.last = nil
	}
	return tmp.data
}

//emptyQueue 判空
func (this *AOVQueue) emptyQueue() bool {
	if this.first == nil && this.last == nil {
		return true
	}
	return false
}
