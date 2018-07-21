package main

import . "fmt"

type I interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}

// receiver是value，按pointer调用go会自动进行转换，不报错
func (s S) Set(age int) {
	s.Age = age
}

func f(i I) {
	i.Set(18)
	Println("i.Get(): ", i.Get())
}

func main() {
	s := S{}
	// receiver是value的method，在内部method对value做出的改变不影响调用者看到的value（如何理解，结果为0？），即按值传递
	f(&s)
	f(s)
}
