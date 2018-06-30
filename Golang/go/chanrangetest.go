package main

import "fmt"

func main() {
    q := make(chan string, 2)
    q <- "one"
    q <- "two"
    close(q)

    for elem1 := range q {
        fmt.Println("elem1 ", elem1)
    }
    for elem := range q {
        fmt.Println(elem)
    }
}
