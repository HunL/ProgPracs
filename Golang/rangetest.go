package main

import "fmt"

func main(){
    list := []string{"a", "b", "c", "d", "e", "f"}
    for k, v := range list{
        fmt.Println("key=", k, "value=", v)
    }

    for pos, char := range "a#x"{
        fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
    }



}
