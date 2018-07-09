package main

import . "fmt"

func main() {
	resInt := romanToInt("IX")
	Println("resInt: ", resInt)
}

func romanToInt(s string) int {
	Println("s: ", s)

	sum := 0
	// for _, c := range s {
	// 	Println("c: ", c)
	// 	switch c {
	// 	case 'I':
	// 		Println("match I")
	// 		sum = sum + 1
	// 	case 'X':
	// 		Println("match X")
	// 		sum = sum + 10
	// 	}
	// }
	strLen := len(s)
	Println("strLen: ", strLen)

	for i := 0; i < strLen; i++ {
		Println("s[i]: ", s[i])
		if s[i] == 'I' {
			nextI := i + 1
			if nextI < strLen {
				if s[nextI] == 'V' {
					Println("nextI equal to 'V'")
					sum += 4
				} else if s[nextI] == 'X' {
					Println("nextI equal to 'V'")
					sum += 9
				}
			} else {
				Println("equal to 'I'")
				sum += 1
			}
		} else if s[i] == 'X' {
			Println("equal to 'X'")
			sum += 10
		}
	}

	return sum
}
