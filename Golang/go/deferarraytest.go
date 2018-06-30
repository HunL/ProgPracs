package main

import (
    "fmt"
    "sort"
)

func main() {
    array := []int{1,2,3,4,5}
    yarra := [5]int{}
    yarra1 := [5]int{}

    for i, v := range array {
        defer fmt.Printf("i: %d, v: %d\n", i, v)
        yarra[i] = v
        //defer func() {yarra1[i] = v}()
        //yarra1[i] = defer func() {return v}()

    }

    fmt.Println("============1")

    for i, v := range yarra {
        fmt.Printf("yarra i: %d, v: %d\n", i, v)
    }
    fmt.Println("============2")

    //sort.Reverse(sort.IntSlice(array))
    sort.Sort(sort.Reverse(sort.IntSlice(array)))
    for i, v := range array {
        fmt.Printf("reverse array i: %d, v: %d\n", i, v)
    }
    fmt.Println("============3")

    for i, v := range yarra1 {
        fmt.Printf("yarra1 i: %d, v: %d\n", i, v)
    }
    fmt.Println("============4")
}
