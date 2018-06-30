package main

import (
	. "fmt"
	"time"
)

func main() {
	time := time.Now()
	Println("time: ", time)

	t := time.Unix()
	Println("t: ", t)
}