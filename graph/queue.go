package main

type QueueNode struct {
	data int
	next *QueueNode
}

type Queue struct {
	first *QueueNode
	last  *QueueNode
}

//NewQueue 初始化队列
func NewQueue() *Queue {
	return &Queue{}
}

//pushQueue 入队
func (this *Queue) pushQueue(data int) {
	tmp := &QueueNode{
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
func (this *Queue) popQueue() int {
	tmp := this.first
	this.first = tmp.next
	if tmp.next == nil {
		this.last = nil
	}
	return tmp.data
}

//emptyQueue 判空
func (this *Queue) emptyQueue() bool {
	if this.first == nil && this.last == nil {
		return true
	}
	return false
}
