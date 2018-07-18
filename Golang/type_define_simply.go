package main

import . "fmt"

// 可以用来简化多参数的函数
type handle func(str string)

func exec(f handle) {
	f("hello")
}

func main() {
	var p = func(str string) {
		Println("first", str)
	}
	exec(p)

	exec(func(str string) {
		Println("second", str)
	})
}
