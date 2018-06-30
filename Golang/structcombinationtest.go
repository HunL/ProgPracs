package main

import . "fmt"

type human struct {
    sex int
}

type person struct {
    human
    name string
    age int
}

type person1 struct {
    human
    name string
    age int
    sex int
}

func main() {
    // p := person{name: "Jack", age: 11}
    // h := human{sex:1}
    p := person{human:human{sex: 1}, name: "Jack", age: 11}


    Println("p: ", p) 
    // Println("h: ", h) 


    p1 := person1{human: human{sex:2}, name: "Tom", age: 18, sex: 0}
    Println("p1.sex: ", p1.sex)
    Println("p1.human.sex: ", p1.human.sex)
}
