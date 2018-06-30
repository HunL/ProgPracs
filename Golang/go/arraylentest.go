package main

import "fmt"

func main() {
    var arr = [][]int {
        {1,2,3},
        {4,5,6},
        {7,8,9},
        {10,11,12,13},
    }
    fmt.Println("len0:",len(arr))
    fmt.Println("len1:",len(arr[3]))
    //fmt.Println("len1:%d",len(arr[][0]))
}
