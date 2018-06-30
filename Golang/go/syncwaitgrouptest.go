package main

import "fmt"

func main(){
    a := 1
    p_a := &a
    go func(p_a *int){
        for i := 1; i < 100000; i++ {
            *p_a += 1
            fmt.Println(*p_a)
        }
    }(p_a)
    go func(p_a *int){
        for i := 1; i < 100000; i++ {
            *p_a -= 1
            fmt.Println(*p_a)
        }
    }(p_a)
}
