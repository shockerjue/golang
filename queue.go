package golang

import (
	"sync"
)

type queueNode struct {
	value interface{}
	next  *queueNode
}

type Queue struct {
	front *queueNode
	rear  *queueNode
	_lock sync.Mutex

	len int
}

// 创建新的链表结构
func NewQueue() *Queue {
	head := &queueNode{
		nil,
		nil,
	}

	queue := &Queue{
		len:   0,
		front: head,
		rear:  head,
	}

	return queue
}

//判断链表是否为空
func (this *Queue) IsEmpty() bool {
	return this.front == this.rear
}

func (this *Queue) Len() int {
	return this.len
}

// 插入值到链表中
func (this *Queue) Push(value interface{}) {
	this._lock.Lock()
	defer this._lock.Unlock()

	node := &queueNode{
		value: value,
		next:  nil,
	}

	rear := this.rear
	for {
		if rear.next == nil {
			break
		}

		rear = rear.next
	}

	rear.next = node
	this.rear = node

	this.len += 1

	return
}

// 返回队列中第一个元素，并将其删除
func (this *Queue) Pop() interface{} {
	empty := this.IsEmpty()
	if empty {
		return nil
	}

	this._lock.Lock()
	defer this._lock.Unlock()

	if nil == this.front.next {
		return nil
	}

	value := this.front.next.value
	this.front = this.front.next

	return value
}

// 遍历链表
func (this *Queue) Foreach(fn func(value interface{})) {
	empty := this.IsEmpty()
	if empty {
		fn(nil)

		return
	}

	node := this.front.next
	for nil != node {
		fn(node.value)

		node = node.next
	}

	return
}
