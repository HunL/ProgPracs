package main

import (
    . "fmt"
    "math"
)

func main() {
    //res := reverse(-120)
    res := reverse(1563847412)
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
    s := 0 
    r := 0
    for x > 0 {
        s = x % 10
        r = r * 10 + s
        x = x / 10 
    }

    maxInt32 := math.Pow(2, 31)
    if float64(r) > maxInt32-1 {
        r = 0
    }

    return r
}
