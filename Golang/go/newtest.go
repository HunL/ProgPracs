package main

import(
    "fmt"
    "sync"
    "bytes"
)

func main(){
    type SyncedBuffer struct{
        lock   sync.Mutex
        buffer bytes.Buffer
    }

    p := new(SyncedBuffer)
    var v SyncedBuffer
    fmt.Printf("%T\n", p)
    fmt.Printf("%v\n", p)
    fmt.Printf("\n", p)
    fmt.Printf("%T\n", v)
    fmt.Printf("%v\n", v)
}
