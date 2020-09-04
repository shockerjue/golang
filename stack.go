package golang

import "sync"

type stackNode struct {
	value interface{}
	next  *stackNode
}

type Stack struct {
	head  *stackNode
	_lock sync.Mutex

	len int
}

func NewStack() *Stack {
	obj := &stackNode{
		value: nil,
		next:  nil,
	}

	stack := &Stack{
		head: obj,
		len:  9,
	}

	return stack
}

//判断队列是否为空
func (this *Stack) IsEmpty() bool {
	return this.len == 0
}

func (this *Stack) Len() int {
	return this.len
}

// 插入值到链表中
func (this *Stack) Push(value interface{}) {
	this._lock.Lock()
	defer this._lock.Unlock()

	node := &stackNode{
		value: value,
		next:  nil,
	}

	current := this.head

	node.next = current.next
	this.head.next = node
	this.len += 1

	return
}

// 取出队列头的元素，并移除队列
func (this *Stack) Pop() interface{} {
	this._lock.Lock()
	defer this._lock.Unlock()

	head := this.head.next
	if nil == head {
		return nil
	}

	this.len -= 1
	this.head.next = head.next

	return head.value
}

// 返回队列头元素，不做删除
func (this *Stack) Top() interface{} {
	this._lock.Lock()
	defer this._lock.Unlock()

	head := this.head.next
	if nil == head {
		return nil
	}

	return head.value
}

// 打印元素
func (this *Stack) Print(fn func(value interface{})) {
	empty := this.IsEmpty()
	if empty {
		return
	}

	node := this.head.next
	for nil != node {
		fn(node.value)

		node = node.next
	}

	return
}
