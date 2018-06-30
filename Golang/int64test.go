package main

import (
	. "fmt"
	"time"
)

func main() {
	t := 30 * time.Second
	Println("time: ", t)

	t1 := int64(t)
	Println("time1: ", t1)

	// dur := time.Duration()
	// Println("dur: ", dur)

	// now := 1529561963
	// then := 1529561933
	now := time.Now()
	Println("now: ", now)
	// then := time.Date(2018, 6, 21, 13, 34, 58, 651387237, time.UTC)
	time.Sleep(3*time.Second)
	then := time.Now()
	Println("then: ", then)

	diff := then.Sub(now)
	Println("diff: ", diff)

	sec := diff.Seconds()
	Println("sec: ", sec)

	if sec >= 3 {
		Println("true")
	}
}