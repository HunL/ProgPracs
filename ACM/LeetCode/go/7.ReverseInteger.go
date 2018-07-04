package main

import . "fmt"

func main() {
    res := reverse(-120)
    Println("res: ", res)
}

func reverse(x int) int {
    helperRes := 0
    if x < 0 {
        helperRes = helper(-x)
        return -helperRes
    }

    helperRes = helper(x)
    return helperRes
}

func helper(x int) int {
    Println("begin x: ", x)
    s := 0 
    r := 0
    for x > 0 {
        s = x % 10
        Println("x: ", x)
        Println("s: ", s)
        r = r * 10 + s
        Println("r: ", r)
        x = x / 10 
    }
 
    return r
}
