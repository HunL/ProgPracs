package main

import (
    . "fmt"
    "sync"
)

var m sync.Map

type User struct {
    ID int
    Name string
}

func main() {
    u1 := User{ID:1, Name:"jack"}
    u2 := User{ID:1, Name:"tom"}

    m.Store(1, u1)
    //Println("m1: ", m)
    m.Range(func(key, value interface{}) bool {
        u := value.(User)
        Println("key: ", key)
        Println("value: ", value)
        Println("u: ", u)
        return true
    })
    m.Store(1, u2)
    //Println("m2: ", m)
    m.Range(func(key, value interface{}) bool {
        u := value.(User)
        Println("key: ", key)
        Println("value: ", value)
        Println("u: ", u)
        return true
    })
}
