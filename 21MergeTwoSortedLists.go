package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//自定义打印链表方法
func (temp *ListNode) myPrint() {
	for temp != nil {
		fmt.Print(strconv.Itoa(temp.Val) + "->")
		temp = temp.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	//头节点设置成-1，就不用提前比较l1.val 和 l2.val了
	nodeHead := new(ListNode)
	nodeHead.Val = -1
	current := nodeHead
	//设置从头部开始遍历
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return nodeHead.Next
}

func main() {

	node4 := new(ListNode)
	node4.Val = 4

	node2 := new(ListNode)
	node2.Val = 2
	node2.Next = node4

	node1 := new(ListNode)
	node1.Val = 1
	node1.Next = node2

	listnode4 := new(ListNode)
	listnode4.Val = 4

	listnode3 := new(ListNode)
	listnode3.Val = 3
	listnode3.Next = listnode4

	listnode1 := new(ListNode)
	listnode1.Val = 1
	listnode1.Next = listnode3

	mergeTwoLists(node1, listnode1).myPrint()
}
