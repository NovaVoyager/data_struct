package main

type QueueNode struct {
	data *Node
	next *QueueNode
}

type Queue struct {
	first *QueueNode
	last  *QueueNode
}

//newQueue 初始化队列
func newQueue() *Queue {
	return &Queue{}
}

//pushQueue 入队
func pushQueue(queue *Queue, data *Node) {
	if data == nil {
		return
	}

	tmp := &QueueNode{
		data: data,
		next: nil,
	}

	if queue.first == nil && queue.last == nil {
		queue.first = tmp
		queue.last = tmp
		return
	}

	queue.last.next = tmp
	queue.last = tmp
}

//popQueue 出队
func popQueue(queue *Queue) *Node {
	tmp := queue.first
	queue.first = tmp.next
	if tmp.next == nil {
		queue.last = nil
	}
	return tmp.data
}

//emptyQueue 判空
func emptyQueue(queue *Queue) bool {
	if queue.first == nil && queue.last == nil {
		return true
	}
	return false
}
