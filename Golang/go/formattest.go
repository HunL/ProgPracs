package main

import "fmt"

type point struct {
    x, y int
}

func main() {
    p := point{1, 2}
    fmt.Println("=======%p, output the value of the point ========")
    fmt.Printf("%p\n", &p)

    fmt.Println("=======%v, output the object value of the struct ========")
    fmt.Printf("%v\n", p)

    fmt.Println("=======%+v, output the object value of the struct ========")
    fmt.Printf("%+v\n", p)


    fmt.Println("=======%#v, output the go format of value ========")
    fmt.Printf("%#v\n", p)

    fmt.Println("=======%T, output the data type of value ========")
    fmt.Printf("%T\n", p)
}
