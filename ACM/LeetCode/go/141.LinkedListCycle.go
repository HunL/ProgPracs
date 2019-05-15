package main

import (
	. "fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	var head = new(ListNode)
	var node1 = new(ListNode)
	var node2 = new(ListNode)
	var node3 = new(ListNode)
	head.Val = 1
	node1.Val = 2
	node2.Val = 3
	node3.Val = 4
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
// 	node3.Next = node2
 	node3.Next = nil
	has := hasCycle(head)
	Print("res: ", has)
}

func hasCycle(head *ListNode) bool {
	cur := head
	for cur != nil {
		Print("cur: ", cur.Val, "\n")
		cur = cur.Next
	}
	return false
} 