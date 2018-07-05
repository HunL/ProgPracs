package main

import . "fmt"

func main() {
	var f = adder()
	Println("f(1): ", f(1))
	Println("f(20): ", f(20))
	Println("f(300): ", f(300))
}

func adder() func(int) int {
	var x int
	return func(delta int) int {
		// Println("\nx: ", x)
		x += delta
		return x
	}
}
