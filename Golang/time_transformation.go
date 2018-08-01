package main

import (
	. "fmt"
	"time"
)

func main() {
	timestamp := time.Now().Unix()
	Println(timestamp)

	t := time.Now()
	Println(t)

	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	Println(tm1)
	Println(tm1.Unix())

	tm2 := tm1.AddDate(0, 0, 1)
	Println(tm2)
	Println(tm2.Unix())
}