package main

import . "fmt"

func main() {
	var p *int
	i := 168
	p = &i
	Println("*p: ", *p)

	// 通过指针p更新i值
	*p = 21
	Println("*p: ", *p)
}
