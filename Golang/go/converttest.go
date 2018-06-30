package main

import "fmt"

func main(){
    mystring := "hello this is string"
    byteslice := []byte(mystring)
    fmt.Printf("%v\n", mystring)
    for k,v := range byteslice{
        fmt.Printf("%v,%v\n", k,v)
    }

    runeslice := []rune(mystring)
    for k,v := range runeslice{
        fmt.Printf("%v,%v\n", k,v)
    }


    b := []byte{'h', 'e', 'l', 'l', 'o'}
    s := string(b)
    i := []rune{257,1024,65}
    r := string(i)
    fmt.Printf("%v\n", b)
    fmt.Printf("%s\n", s)
    fmt.Printf("%v\n", i)
    fmt.Printf("%s\n", r)
}
