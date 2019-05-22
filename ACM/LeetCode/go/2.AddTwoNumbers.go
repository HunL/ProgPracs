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

    // num := reverse_num(123, 0)
    // Print("reverse num: ", num)


    // test1 := getListNode1(1723)
    // test1 := getListNode1(173)
    // 第一位进位的处理
    // 最后一位进位的处理
    // 第2个数比第1个数长
    // 0的处理
    test1 := getListNode1(768)
    // test1 := getListNode1(0)
    showNodes(test1)
    
    // test2 := getListNode1(3435)
    // test2 := getListNode1(345)
    // test2 := getListNode1(6455)
    test2 := getListNode1(0)
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

// @todo 画画写写加深理解
func reverse_node(node *ListNode) *ListNode {
    // if (node == nil) {
    //     return node
    // }
    
    // a := num % 10
    // b := num / 10
    // result = result * 10 + a
    // return reverse_num(b, result)
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

// 先写好所有标准测试用例
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var l3 = new(ListNode)
    Print("\nfirst 1 2: ", l1.Val, "\t", l2.Val, "\n")
    l3.Val = (l1.Val+l2.Val) % 10
    // l1 = reverse_node(l1)
    // l2 = reverse_node(l2)
    remainTmp := (l1.Val+l2.Val) / 10
    l1 = l1.Next
    l2 = l2.Next
    var valTmp int
    
    // for (l1 != nil && remainTmp != 0) {
    // for (l1 != nil) || (remainTmp != 0) {
    for (l1 != nil) || (l2 != nil) || (remainTmp != 0)  {
        
        var tail = new(ListNode)
        if (l1 != nil) {
            if (l2 != nil) {
                valTmp = l1.Val+l2.Val+remainTmp
                Print("\nadd 1 2: ", l1.Val, "\t", l2.Val, "\t", remainTmp)
            } else {
                valTmp = l1.Val+remainTmp
                Print("\nadd 1: ", l1.Val, "\t", remainTmp)
            }
        } else {
            if (l2 != nil) {
                valTmp = l2.Val+remainTmp
                Print("\nadd 2: ", l2.Val, "\t", remainTmp)
            } else {
                valTmp = remainTmp
                Print("\nadd: ", remainTmp)
            }
        }
        
        Print("\nadd: ", valTmp)
        
        if (valTmp >= 10) {
            tail.Val = valTmp % 10
            remainTmp = valTmp / 10
        } else {
            tail.Val = valTmp
            remainTmp = valTmp / 10
        }    
        
        if (l1 != nil) {
            l1 = l1.Next   
        }
        if (l2 != nil) {
            l2 = l2.Next
        }

        // 这个地方想了一个小时！！！大爷的
        tail.Next = l3.Next
        l3.Next = tail
        // l3 = l3.Next
        
        Println()
    }

    // if (remainTmp != 0) {
    //     Print("\nadd: ", remainTmp)
    //     var tail = new(ListNode)
    //     valTmp := remainTmp
    //     Print("\nadd: ", valTmp)
        
    //     if (valTmp >= 10) {
    //         tail.Val = valTmp % 10
    //         remainTmp = valTmp / 10
    //     } else {
    //         tail.Val = valTmp
    //     }    
        
    //     tail.Next = l3.Next
    //     l3.Next = tail
    //     // l3 = l3.Next
        
    //     Println()
    // }

    l3.Next = reverse_node(l3.Next)
    return reverse_node(l3)
}