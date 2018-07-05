package main

import . "fmt"

func main() {
    //res := isPalindrome(121)
    res := isPalindrome(-121)
    Println("res: ", res)
}

func isPalindrome(x int) bool {

    if x < 0 {
        return false
    }

    beginx := x
    Println("begin x: ", x)

    s := 0
    r := 0
    for x > 0 {
        s = x % 10
        r = r * 10 + s
        x = x / 10
    }
    Println("r: ", r)

    isEqual := false
    if r == beginx {
        Println("r is equal to x")
        isEqual = true
    }
    return isEqual
}
