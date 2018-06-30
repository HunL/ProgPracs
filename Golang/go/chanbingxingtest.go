package main

import "fmt"

func main() {
    //buf()
    unbuf()
}


func buf() {
    message := make(chan string)

    go func(){
        message <-"ping"
    }()

    fmt.Println(<-message)
}

func unbuf() {
    message := make(chan string, 2)

    go func(){
        message <-"ping"
        message <-"pong"
    }()

    //fmt.Println("len ", len(message))

    close(message)
    // cannot for unbuf chan?
    s, v := <-message
    fmt.Println("get value is %d, %d\n", s, v)

    for e := range <-message {
        fmt.Println(e)
    }
}
