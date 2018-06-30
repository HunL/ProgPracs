package main

import "fmt"

func unhex(c byte) byte{
    
    switch{
        case '0' <= c && c <= '9':
            return c - '0'
        case 'a' <= c && c <= 'f':
            return c - 'a' + 10
        case 'A' <= c && c <= 'F':
            return c - 'A' + 10
        default:
            return 99
    }
    return 0
}

func main(){
    var c,d byte
//    c = 'E'
    d = unhex(c)
    fmt.Println(d)
}
