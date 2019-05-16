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
	node3.Next = node2
 	// node3.Next = nil
	has := hasCycle(head)
	Print("res: ", has)
}

func hasCycle(head *ListNode) bool {
	cur := head
	var arr []*ListNode
	flag := false
	for cur != nil  {
		Print("cur: ", cur.Val, "\n")
		for k,v := range arr {
			Print("arr i: ", k, v, "\n")
			// Print("arr i: ", i.Val, "\n")
			if (v == cur) {
				flag = true
				return true
			}
		}
		if (flag == false) {
			arr = append(arr, cur)
		}
		cur = cur.Next
	}
	return false
} 