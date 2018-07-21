package main

import . "fmt"

// 多个类型实现了interface I，这些类型的值都可以直接使用interface I的变量存储
type I interface {
	get() int
	set()
}

type S struct {
	Age int
}

func (s S) get() int {
	// Println("get: ", s.Age)
	return s.Age
}

func (s *S) set() {
	s.Age = 18
	// Println("set: ", s.Age)
}

type M map[string]int

func (m M) get() int {
	// Println("get: ", m["test"])
	return m["test"]
}

func (m M) set() {
	m["test"] = 3
	// Println("set: ", m.get())
}

func f(i I) {
	i.set()
	Println(i.get())
}

func main() {
	s := S{}
	// (&s).set()
	// i := &s
	// i.set()
	f(&s)

	m := make(M)
	// (&m).set()
	f(&m)
}
