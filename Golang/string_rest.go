package main

import . "fmt"

func main() {
	str := "abcdefg"
	Println("str: ", str)

	len := len(str)
	Println("len: ", len)

	str0 := str[0:len]
	Println("str0: ", str0)

	str1 := str[1:len]
	Println("str1: ", str1)

	str2 := str[2:len]
	Println("str2: ", str2)
}
