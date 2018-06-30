package main

import "fmt"

func main(){
    var i, j, times, div, remain int
    i = 25 
    j = 10
    times = i * j
    fmt.Println("hello, i =", i, "j =", j, "times =",times)

    div = i / j
    fmt.Println("div =", div)

    remain = i % j
    fmt.Println("remain =", remain)
}
