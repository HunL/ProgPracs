package main

import . "fmt"


func test1() {
    var i interface{} = "99"
    j, b := i.(int)
    if b {
        Println("%T->%d\n", i, j)
    } else {
        Println("type transfer error")
    }
}

func test2() {
    var i interface{} = 99
    j := i.(int)
    Println("%T->%d\n", i, j)
}

func test3() {
    var i1 int = 7
    Printf("%T->%v\n", i1, i1)
    i2 := float32(i1)
    i3 := int64(i1)
    Printf("%T->%v\n", i2, i2)
    Printf("%T->%v\n", i3, i3)
}

func test4() {
    var i1 int = 7
    Printf("%T->%v\n", i1, i1)
    i2 := int8(i1)
    //i3 := []int8(i1)
    //i4 := []string(i1)
    Printf("%T->%v\n", i2, i2)
    //Printf("%T->%v\n", i3, i3)
    //Printf("%T->%v\n", i4, i4)
}


func main() {
    //test1()
    //test2()
    //test3()
    test4()
}
