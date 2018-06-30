package main

import "fmt"

func main(){
    s0 := []int{0, 0}
    s1 := append(s0, 2)
    s2 := append(s1, 3,5,7)
    s3 := append(s2, s0...)

    for k,v := range s0{
        fmt.Println(k,v)
    }
    fmt.Println("\n")
    for k,v := range s1{
        fmt.Println(k,v)
    }
    fmt.Println("\n")
    for k,v := range s2{
        fmt.Println(k,v)
    }
    fmt.Println("\n")
    for k,v := range s3{
        fmt.Println(k,v)
    }
}
