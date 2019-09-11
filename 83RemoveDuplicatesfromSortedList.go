package main

import "fmt"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head

	for q := head.Next; q != nil; q = q.Next {
		if q.Val != p.Val {
			p.Next = q
			p = q
		}
	}
	p.Next = nil
	return head
}

func main() {
	fmt.Println("")
}
