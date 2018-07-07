package main

import . "fmt"

type vertex struct {
	x int
	y int
}

func main() {
	v := vertex{1, 1}
	p := &v
	Println("p.x: ", p.x)
	Println("v: ", v)
	Println("*p: ", *p)

	p.x = 1e9
	Println("p.x: ", p.x)
	Println("v: ", v)
	Println("*p: ", *p)
}
