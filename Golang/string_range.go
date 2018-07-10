package main

import . "fmt"

func main() {
	str := "abcdefg"
	for k, v := range str {
		Println("k: ", k)
		Println("v: ", v)
	}
}
