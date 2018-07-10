package main

import . "fmt"

func main() {
	str := "[{}]"
	isValid := isValid(str)
	Println("isValid: ", isValid)
}

func isValid(s string) bool {
	Println("s: ", s)

	return false
}
