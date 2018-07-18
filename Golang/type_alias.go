package main

import . "fmt"

type a = string
type b string

func sayA(str a) {
	Println(str)
}

func sayB(str b) {
	Println(str)
}

func main() {
	var s = "test"
	sayA(s)
	// sayB(s)
}
