package main

import (
	. "fmt"
	"strings"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	// strs := []string{}
	retPrefix := longestCommonPrefix(strs)
	Println("retPrefix: ", retPrefix)
}

func longestCommonPrefix(strs []string) string {
	Println("strs: ", strs)
	if len(strs) == 0 {
		return ""
	}

	str0 := strs[0]
	Println("str0: ", str0)

	commonPrefix := ""
	Println("init commonPrefix: ", commonPrefix)

Loop:
	for k0, v0 := range str0 {
		Println("k0: ", k0)
		Println("v0: ", v0)

		Println("string(v0): ", string(v0))
		commonPrefix += string(v0)
		Println("commonPrefix: ", commonPrefix)

		for k, v := range strs {
			Println("k: ", k)
			Println("v: ", v)
			if strings.HasPrefix(v, commonPrefix) == false {
				Println("doesn't has prefix v: ", v)
				Println("commonPrefix: ", commonPrefix)
				commonPrefix = commonPrefix[0 : len(commonPrefix)-1]
				break Loop
			}
		}
	}

	return commonPrefix
}
