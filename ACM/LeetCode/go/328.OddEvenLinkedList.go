package main

import (
    . "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {    
    test1 := getListNode(768)
    showNodes(test1)
    
    test2 := oddEvenList(test1)
    showNodes(test2)
}

// 先写好所有标准测试用例
func oddEvenList(head *ListNode) *ListNode {
    Print("oddEvenList\n")
    return head
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

func showNodes(p *ListNode) {
    for (p != nil) {
        Print("show: ", p.Val, "\t")
        p = p.Next
    }
    Println()
}
