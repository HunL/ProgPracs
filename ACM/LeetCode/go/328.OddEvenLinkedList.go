package main

import (
    . "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {    
    test1 := getListNode(75816)
    showNodes(test1)
    
    test2 := oddEvenList(test1)
    showNodes(test2)
}

func oddEvenList(node *ListNode) *ListNode {
    Print("oddEvenList\n")
    index := 1
    cur := node
    // var pre *ListNode = nil
    pre := node
    for (cur != nil) {
       index += 1 
       Print("index: ", index, "\n")
       if (index % 2 == 0) {
           if (cur.Next != nil) {
               pretmp := cur.Next
               curtmp := cur
               
               tmp := cur
               pre.Next = cur.Next
               cur.Next = cur.Next.Next
               pre.Next.Next = tmp
               
               pre = pretmp
               cur = curtmp
           }
        //   Print("odd % 2: ", odd % 2, "\n")    
       } else {
           pre = cur
           cur = cur.Next    
       }
       
    }
    return node
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