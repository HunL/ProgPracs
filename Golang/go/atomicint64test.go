package main

import (
    "fmt"
    "sync/atomic"
)

func main() {
    var i64 int64
    fmt.Println("=====old i64 value=====")
    fmt.Println(i64)

    newI64 := atomic.AddInt64(&i64, 3)

    fmt.Println("=====new i64 value=====")
    fmt.Println(i64)
    fmt.Println(newI64)
}
