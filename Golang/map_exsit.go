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

var s = map[int][]int{
	1: {1, 2, 3},
	2: {2, 3, 4},
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
	Println("array test ====================")
	Println("s[1]: ", s[1])
	Println("s[2]: ", s[2])
	Println("s[4]: ", s[4])
}
