package main

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := new(Stack)
	stack.Push(3, 2, 4, 1, 5)
	fmt.Println(stack)
	fmt.Println(stack.Pop(), stack)
}
