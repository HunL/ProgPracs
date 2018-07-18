package main

import . "fmt"

type S string

func (s *S) Hi() {
	Println("S hi")
}

type T = S

// 类型别名增加的方法，原类型也可以使用
func (t *T) Hello() {
	Println("T hello")
}

func exec(obj *S) {
	obj.Hi()
	obj.Hello()
}

func main() {
	s := new(S)
	t := new(T)
	exec(s)
	exec(t)
}
