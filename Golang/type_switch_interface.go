package main

import . "fmt"

func main() {
	// var a interface{} = "abc"
	// var a interface{} = 12
	var a interface{} = []int{1, 2, 3}

	switch v := a.(type) {
	case string:
		Println("string type")
	case int:
		Println("int type")
	default:
		Println("other type: ", v)
	}
}
