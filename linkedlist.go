package golang

import (
	"errors"
	"sync"
)

type ListNode struct {
	value interface{}
	next  *ListNode
}

type LinkedList struct {
	head  *ListNode
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
		head: head,
	}

	return linked
}

//判断链表是否为空
func (this *LinkedList) IsEmpty() bool {
	return this.len == 0
}

// 获取链表的长度
func (this *LinkedList) Len() int {
	return this.len
}

// 插入值到链表中
// @param value
func (this *LinkedList) Insert(value interface{}) {
	this._lock.Lock()
	defer this._lock.Unlock()

	node := &ListNode{
		value: value,
		next:  nil,
	}

	current := this.head
	for {
		if current.next == nil {
			break
		}

		current = current.next
	}

	current.next = node
	this.len += 1

	return
}

// 删除链表对应的值
// 从链表中移除值为value的项
func (this *LinkedList) Remove(value interface{}) error {
	empty := this.IsEmpty()
	if empty {
		return errors.New("This is an empty list")
	}

	this._lock.Lock()
	defer this._lock.Unlock()

	current := this.head
	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			this.len -= 1

			return nil
		}

		current = current.next
	}

	return nil
}

// 遍历链表
// @param fn 迭代调用函数，参数为链表中的值
func (this *LinkedList) Foreach(fn func(value interface{}) bool) {
restart:

	empty := this.IsEmpty()
	if empty {
		return
	}

	node := this.head.next
	for nil != node {
		if fn(node.value) {
			this.Remove(node.value)

			goto restart
		}

		node = node.next
	}

	return
}

// 打印链表中元素
// @param fn 打印链表中的元素
func (this *LinkedList) Print(fn func(value interface{})) {
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
