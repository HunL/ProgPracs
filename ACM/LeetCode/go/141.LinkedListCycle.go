package main

import (
	. "fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	var head *ListNode
	var node1 *ListNode
	var node2 *ListNode
	head.Val = 1
	node1.Val = 2
	node2.Val = 3
	head.Next = node1
	node1.Next = node2
	node2.Next = nil
	has := hasCycle(head)
	Print("res: ", has)
}

func hasCycle(head *ListNode) bool {
	cur := head
	for cur != nil {
		Print("cur: ", cur.Val)
		cur = cur.Next
	}
	return false
}