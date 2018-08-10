package main

import . "fmt"

func main() {
	Println("aaa")
	a := 1
	Println("a: ", a)

	arr := []int{1, 2, 3}
	len := len(arr)
	Println("len: ", len)

	// 写 arr1 := arr[2:len(arr)] 报错，无法测试 gofmt -r 'a[n:len(a)] -> a[n:]' -w gofmttest.go
	// arr1 := arr[2:len(arr)]
	// 写 arr1 := arr[2:len]，可以测试 gofmt -r 'a[n:len] -> a[n:]' -w gofmttest.go
	arr1 := arr[2:]
	Println("arr1[0]: ", arr1[0])

	// b := (a)，测试gofmt -r '(a) -> a' -w gofmttest.go
	b := a
	Println("b: ", b)
}
