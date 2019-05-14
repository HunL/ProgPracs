package main

import (
    . "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    // test(12)
    // showNodes(test1)

    test1 := getListNode1(123)
    showNodes(test1)

    test2 := getListNode1(345)
    showNodes(test2)

    test3 := addTwoNumbers(test1, test2)
    showNodes(test3)

    // l1 := getListNode(123)
    // showNodes(l1)

    // l2 := getListNode(345)
    // showNodes(l2)

    // l3 := addTwoNumbers(l1, l2)
    // showNodes(l3)
}
func test(num int) {
    var val = num % 10
    Print("test val: ", val)

    var val2 = num / 10
    Print("\t test val2: ", val2, "\n")
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

func getListNode1(num int) *ListNode {
    Print("\norigin num: ", num)
    var head = new(ListNode)
    if (num == 0) {
        head.Val = 0
        return head
    }
    head.Val = num % 10
    num = num / 10
    // num = reverse_num(num)
    for (num != 0) {
        Print("\nnum: ", num, "\t")
        var node = new(ListNode)
        node.Val = num % 10
        Print("\nnode.Val: ", node.Val, "\t")
        num = num / 10
        node.Next = head.Next
        head.Next = node
    }
    Println()
    return head
}

func reverse_num(num int) int {
    if (num == 0) {
        return num
    }
    
    for (num != 0) {
        // a := num % 10
        // b := num / 10

    }
    return 0
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
    l3.Val = l1.Val+l2.Val
    l1 = l1.Next
    l2 = l2.Next
    
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