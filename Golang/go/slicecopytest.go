package main

import "fmt"

func main(){
    var a = [...]int{0,1,2,3,4,5,6,7}
    var s = make([]int, 6)

    n1 := copy(s, a[0:])
    for k,v := range s{
        fmt.Println(k,v)
    }
    fmt.Printf("\n")

    n2 := copy(s, s[2:])

    fmt.Println(n1)
    fmt.Println(n2)
    fmt.Printf("\n")

    for k,v := range s{
        fmt.Println(k,v)
    }
}
