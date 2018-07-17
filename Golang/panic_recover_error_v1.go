package main

import (
	"errors"
	. "fmt"
)

func main() {
	err := a()
	if err != nil {
		Println("err: ", err)
	}
}

func a() (err error) {
	err = b()
	if err != nil {
		return
	}
	err = c()
	if err != nil {
		return
	}
	err = errors.New("error inside a")
	return err
}

func b() (err error) {
	err = errors.New("error inside b")
	return
}

func c() (err error) {
	err = errors.New("error inside c")
	return
}
