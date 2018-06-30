package main

import "fmt"

type NameAge struct{
    name string
    age int
}

func main(){
    a := new(NameAge)
    fmt.Printf("%v\n", a)

    a.name = "Pete"; a.age = 42
    fmt.Printf("%v\n", a)

    fmt.Printf("%s\n", a.name)
    fmt.Printf("%d\n", a.age)
}
