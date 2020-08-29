package golang

import (
	"fmt"
	"testing"
	"time"
)

func TestLinkedList(t *testing.T) {
	linkedlist := NewLinkedList()
	linkedlist.Insert(12)
	linkedlist.Insert(15)
	linkedlist.Insert(11)
	linkedlist.Insert(13)
	linkedlist.Insert(100)
	linkedlist.Insert(121)
	linkedlist.Insert(111)

	go func() {
		for i := 0; i < 10; i++ {
			linkedlist.Insert(11 * (i + 22))
		}
	}()

	go func() {
		linkedlist.Foreach(func(value interface{}) bool {
			if 12 == value.(int) {
				fmt.Println("Remove 12 success!")

				return true
			}

			if 111 == value.(int) {
				fmt.Println("Remove 341 success!")

				return true
			}

			return false
		})

	}()

	time.Sleep(2 * time.Second)

	linkedlist.Print(func(value interface{}) {
		fmt.Println(value)
	})

	return
}
