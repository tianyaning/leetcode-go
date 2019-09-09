package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func isValid(s string) bool {
	stack := NewStack()
	for _, c := range s {
		if c == '(' {
			stack.Push(')')
		} else if c == '{' {
			stack.Push('}')
		} else if c == '[' {
			stack.Push(']')
		} else if stack.Len() == 0 || stack.Pop() != c {
			return false
		}
	}

	return stack.Len() == 0
}

func main() {
	fmt.Print(isValid("()"))
	fmt.Print(isValid("()[]{}"))
	fmt.Print(isValid("(]"))
	fmt.Print(isValid("([)]"))
	fmt.Print(isValid("{[]}"))
}
