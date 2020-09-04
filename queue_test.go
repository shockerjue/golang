package golang

import (
	"fmt"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	fmt.Println(queue.Pop())

	queue.Push(1)
	queue.Push(2)

	go func() {
		defer func() {
			fmt.Println("go 1 over!")
		}()
		for i := 0; i < 3; i++ {
			time.Sleep(time.Duration(1) * time.Second)
			queue.Push(i * 3)
			queue.Push((i + 1) * 4)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Duration(1) * time.Second)
			queue.Pop()
		}
	}()

	time.Sleep(time.Duration(5) * time.Second)
	queue.Foreach(func(value interface{}) {
		fmt.Println(value)
	})

	return
}
