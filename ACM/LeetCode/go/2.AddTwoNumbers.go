package main

import (
    . "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    l1 := getListNode(123)
    showNodes(l1)

    l2 := getListNode(345)
    showNodes(l2)

    l3 := addTwoNumbers(l1, l2)
    showNodes(l3)
}

func getListNode(num int) *ListNode {
    var head = new(ListNode)
    head.Val = num % 10
    var node1 = new(ListNode)
    node1.Val = (num  % 100) / 10
    var node2 = new(ListNode)
    node2.Val = (num  % 1000) / 100

    head.Next = node1
    node1.Next = node2
    return head
}

func showNodes(p *ListNode) {
    for (p != nil) {
        Print("show: ", p.Val, "\t")
        p = p.Next
    }
    Println()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3 = new(ListNode)
    
    for (l1 != nil) {
        Print("add: ", l1.Val, "\t", l2.Val)
        var tail = new(ListNode)
        tail.Val = l1.Val+l2.Val
        l1 = l1.Next
        l2 = l2.Next
        // 这个地方想了一个小时！！！大爷的
        tail.Next = l3.Next
        l3.Next = tail
        // l3 = l3.Next
        
        Println()
    }
    return l3
}