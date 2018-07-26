package main

import . "fmt"

var m = map[int]bool{
	// 1:true,
	2: true,
	3: true,
}

var i = map[int]int{
	// 1:11,
	2: 22,
	3: 33,
}

func main() {
	Println("bool test ====================")
	Println("m[1]: ", m[1])
	Println("m[2]: ", m[2])
	Println("m[4]: ", m[4])
	Println("int test ====================")
	Println("i[1]: ", i[1])
	Println("i[2]: ", i[2])
	Println("i[4]: ", i[4])
}
