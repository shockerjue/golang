package golang

import (
	"errors"
	"sync"
)

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

type LinkedList struct {
	Head  *ListNode
	_lock sync.Mutex

	len int
}

// 创建新的链表结构
func NewLinkedList() *LinkedList {
	head := &ListNode{
		nil,
		nil,
	}

	linked := &LinkedList{
		len:  0,
		Head: head,
	}

	return linked
}

//判断链表是否为空
func (this *LinkedList) IsEmpty() bool {
	return this.len == 0
}

func (this *LinkedList) Len() int {
	return this.len
}

// 插入值到链表中
func (this *LinkedList) Insert(value interface{}) {
	this._lock.Lock()
	defer this._lock.Unlock()

	node := &ListNode{
		Value: value,
		Next:  nil,
	}

	current := this.Head
	for {
		if current.Next == nil {
			break
		}

		current = current.Next
	}

	current.Next = node
	this.len += 1

	return
}

// 删除链表对应的值
func (this *LinkedList) Remove(value interface{}) error {
	empty := this.IsEmpty()
	if empty {
		return errors.New("This is an empty list")
	}

	this._lock.Lock()
	defer this._lock.Unlock()

	current := this.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			this.len -= 1

			return nil
		}

		current = current.Next
	}

	return nil
}

// 遍历链表
func (this *LinkedList) Foreach(fn func(value interface{}) bool) {
restart:

	empty := this.IsEmpty()
	if empty {
		return
	}

	node := this.Head.Next
	for nil != node {
		if fn(node.Value) {
			this.Remove(node.Value)

			goto restart
		}

		node = node.Next
	}

	return
}

// 打印链表中元素
func (this *LinkedList) Print(fn func(value interface{})) {
	empty := this.IsEmpty()
	if empty {
		return
	}

	node := this.Head.Next
	for nil != node {
		fn(node.Value)

		node = node.Next
	}

	return
}
