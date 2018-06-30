package main

import "fmt"


func main() {
    a()
    fmt.Println("===============")
    bint := b()
    fmt.Printf("bint = %d\n", bint)
    fmt.Println("===============")
    c()
    fmt.Println("===============")
    d()
}

func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}

func b() (i int) {
    defer func() {i++}()
    return 1
}

func c() {
    defer fmt.Printf("1\n")
    defer fmt.Printf("2\n")
    defer fmt.Printf("3\n")
    defer fmt.Printf("4\n")
}

func d() {
    a := 1
    b := 2
    defer calc("1", a, calc("10", a, b))
    a = 0
    return
}

func calc(index string, a, b int) int {
    ret := a + b
    fmt.Println(index, a, b, ret)
    return ret
}
