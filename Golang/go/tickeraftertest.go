package main

import (
    . "fmt"
    "time"
)

func f() *time.Timer {
    return time.AfterFunc(2*time.Second, func() {
        Println("expired")
    })
}

func main() {
    t := f()
    Println("t: ", t)
    t.Stop()
    time.Sleep(4*time.Second)
}   
