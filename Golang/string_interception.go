package main

import . "fmt"

func main() {
	s := "abcdefg"
	Println("s: ", s)

	s = string([]rune(s)[:3])
	Println("s: ", s)

	zhs := "ab我中文cdefd"
	Println("zhs: ", zhs)

	zhs = string([]rune(zhs)[:4])
	Println("zhs: ", zhs)
}
