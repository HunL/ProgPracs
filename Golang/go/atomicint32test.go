package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var i32 int32
    fmt.Println("=====old i32 value=====")
    fmt.Println(i32)

    newI32 := atomic.AddInt32(&i32, 3)

    fmt.Println("=====new i32 value=====")
    fmt.Println(i32)
    fmt.Println(newI32)
}
