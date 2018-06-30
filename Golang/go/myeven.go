package main

import(
    "even"
    "fmt"
)

func main(){
    i := 6
    fmt.Printf("Is %d even? %v\n", i, even.odd(i))
}
