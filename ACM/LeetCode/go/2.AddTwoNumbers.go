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
    head.Val = 1
    var node1 = new(ListNode)
    node1.Val = 2
    var node2 = new(ListNode)
    node2.Val = 3

    head.Next = node1
    node1.Next = node2
    showNodes(head)
}

func showNodes(p *ListNode) {
    for (p != nil) {
        Println(p.Val)
        p = p.Next
    }
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
 
    return l1
}