package main

import "fmt"

type Person struct{
    Name string
    Age  int
}

func (p *Person) sayHi1() {
    p.Name = "jeff1"
}

func (p Person) sayHi2() {
    p.Name = "jeff2"
}

func main() {
    p1 := &Person{Name:"test", Age:10}
    fmt.Println("name1:", p1.Name)
    p1.sayHi1()
    fmt.Println("name2:", p1.Name)

    p2 := &Person{Name:"test1", Age:11}
    fmt.Println("name3:", p2.Name)
    p2.sayHi2()
    fmt.Println("name4:", p2.Name)

}
