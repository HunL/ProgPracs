package main

import (
	. "fmt"
	"time"
)

func main() {
	time1 := time.Now()
	Println("time1: ", time1)
	time.Sleep(900 * time.Millisecond)

	time2 := time.Now()
	Println("time2: ", time2)

	diff := time2.Sub(time1)
	Println("diff: ", diff)

	if diff >= 3 {
		Println("bigger")
	} else {
		Println("smaller")
	}

	diffSec := diff.Seconds()
	Println("diffSec: ", diffSec)
	if diffSec >= 3 {
		Println("sec bigger")
	} else {
		Println("sec smaller")
	}
}