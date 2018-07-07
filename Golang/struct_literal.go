package main

import . "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 1}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{3, 3}
)

func main() {
	Println("v1: ", v1)
	Println("v2: ", v2)
	Println("v3: ", v3)
	Println("p: ", p)
}
