package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    array := []int{1,2,3,4}
    for i := 0; i < 4; i++ {
        r := rand.New(rand.NewSource(time.Now().UnixNano()))
        from := r.Intn(len(array))
        fmt.Println("from: ", from)
    }
}
