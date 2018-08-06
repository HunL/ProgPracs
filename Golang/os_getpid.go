package main

import (
	. "fmt"
	"os"
)

func test1() {
	Println("test1 pid: ", os.Getpid())
}

func test2() {
	Println("test2 pid: ", os.Getpid())
}

func main() {
	test1()
	test2()
	Println("main pid: ", os.Getpid())
}