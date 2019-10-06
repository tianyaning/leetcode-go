package main

import (
	"fmt"
)

func removeDuplicates(S string) string {
	stack := NewStack()
	for _, value := range S {
		if stack.Peak() == value {
			stack.Pop()
		} else {
			stack.Push(value)
		}
	}

	result := ""
	list := stack.list

	for e := list.Front(); e != nil; e = e.Next() {
		//fmt.Println("type:", reflect.TypeOf(e.Value))
		switch value := e.Value.(type) {
		case int32:
			result = result + string(value)
		}
	}

	return result
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))

}
