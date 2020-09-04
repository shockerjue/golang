package golang

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	fmt.Println(stack.Pop())
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	stack.Print(func(value interface{}) {
		fmt.Println(value)
	})

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println("Pop two element:")
	stack.Print(func(value interface{}) {
		fmt.Println(value)
	})

	return
}
