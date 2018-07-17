package main

import (
	. "fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			Printf("Runtime error caught: %v\n", r)
		}
	}()

	a()
}

func a() {
	b()
	c()

	panic("error inside a")
}

func b() {
	panic("error inside b")
}

func c() {
	panic("error inside c")
}
