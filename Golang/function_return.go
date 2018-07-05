package main

import . "fmt"

func main() {
	f := add()
	Println("f(2): ", f(2))
	f2 := add2(3)
	Println("f2(2): ", f2(2))
}

func add() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func add2(a int) func(b int) int {
	return func(b int) int {
		Println("add2 a: ", a, ", b: ", b)
		return a + b
	}
}
