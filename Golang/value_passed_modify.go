package main

import . "fmt"

func main() {
	i := 10
	Println("before modify i: ", i)
	modify(&i)
	Println("after modify i: ", i)
}

func modify(addr *int) {
	*addr = 20
}
