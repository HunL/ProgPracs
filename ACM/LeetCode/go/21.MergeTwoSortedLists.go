package main

import (
    . "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    l1 := getListNode(136)
    showNodes(l1)
    
    l2 := getListNode(258)
    showNodes(l2)

    l3 := mergeTwoLists(l1, l2)
    showNodes(l3)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    // cur1 := l1
    // cur2 := l2
    res := l1
    // for (cur1 != nil) {
    //     for (cur2 != nil) {
    for (l1 != nil) {
        Print("a: ", l1.Val, "\t")
        // Print("a2: ", l2.Val, "\n")
        for (l2 != nil) {
            if (l1.Val <= l2.Val) {
                Print("b: ", l1.Val, "\t")
                // Print("b2: ", l2.Val, "\n")
                Print("b2: ", l2.Val, "\t")
                res.Next = l1.Next
                Print("b3: ", res.Next.Val, "\n")
                // l2 = l2.Next
            } else {
                Print("c: ", l1.Val, "\t")
                Print("c2: ", l2.Val, "\n")
                res.Next = l2
                Print("c3: ", res.Next.Val, "\n")
                // l2 = l2.Next
            }
            l2 = l2.Next
        }
        Print("fuck l1: ", l1.Val, "\n")
        l1 = l1.Next
    }
    return res
}


func getListNode(num int) *ListNode {
    Print("\norigin num: ", num)
    var head = new(ListNode)
    if (num == 0) {
        head.Val = 0
        return head
    }
    
    head.Val = num % 10
    num = num / 10
    num = reverse_num(num, 0)
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
    head = reverse_node(head)
    return head
}

func reverse_num(num int, result int) int {
    if (num == 0) {
        return result
    }
    
    a := num % 10
    b := num / 10
    result = result * 10 + a
    return reverse_num(b, result)
}

func reverse_node(node *ListNode) *ListNode {
    cur := node
    var pre *ListNode = nil
    for cur != nil {
        pre, cur, cur.Next = cur, cur.Next, pre
    }
    return pre
}

func showNodes(p *ListNode) {
    for (p != nil) {
        Print("show: ", p.Val, "\t")
        p = p.Next
    }
    Println()
}