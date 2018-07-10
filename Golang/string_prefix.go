package main

import (
	. "fmt"
	"strings"
)

func main() {
	str := "abcdefg"
	hasPrefix := strings.HasPrefix(str, "abc")
	Println("hasPrefix: ", hasPrefix)

	hasVoidPrefix := strings.HasPrefix(str, "")
	Println("hasVoidPrefix: ", hasVoidPrefix)
}
