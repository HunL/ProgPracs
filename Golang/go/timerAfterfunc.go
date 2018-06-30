package main

import (
    "fmt"
    "time"
)

func main() {
    var t *time.Timer

    f := func() {
        fmt.Printf("Expiration time: %v.\n", time.Now())
        fmt.Printf("C's len: %d\n", len(t.C))
    }

    t = time.AfterFunc(3 * time.Second, f)
    time.Sleep(5 * time.Second)
}
