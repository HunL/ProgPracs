package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    i := r.Intn(2)
    fmt.Println("i: ", i)
}
