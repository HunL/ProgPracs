package main

import (
    "fmt"
//    "bytes"
)

func main() {
    a1 := []int{1,4,6}
    a2 := []int{5,8,2}
    a := append(a1,a2...)
    fmt.Println("a: ", a)
}
