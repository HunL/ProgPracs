package main

import "fmt"

func main() {
    a := func(){
        fmt.Println("Hello")
    }
    fmt.Printf("%T\n", a)
    a()
}
