package main

import "fmt"

/*
type I interface{
    Get() int
    Put(int)
}
*/

type S struct {i int}
func (p *S) Get() int {return p.i}
func (p *S) Put(v int) {p.i = v}


func main(){
    var s S
//    p := &s
//    f(p)
    f(s)
}

//func f(p I){
func f(p S){
    fmt.Println(p.Get())
    p.Put(1)
    fmt.Println(p.Get())
    p.Put(3)
    fmt.Println(p.Get())
}
