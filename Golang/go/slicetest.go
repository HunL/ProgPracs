package main

import "fmt"

func main(){
    var array [100]int
    slice := array[0:100]

    slice[98] = 'a'
    slice[99] = 'c'

    for k,v := range slice{
        fmt.Println(k,v)
    }

}
