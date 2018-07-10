package main

import . "fmt"

func main() {
	resInt := romanToInt("MCMXCIV", 0)
	Println("resInt: ", resInt)
}

func romanToInt(s string, sum int) int {
	Println("\ns: ", s)
	Println("sum: ", sum)

	strLen := len(s)
	Println("strLen: ", strLen)

	if strLen == 0 {
		Println("retSum: ", sum)
		return sum
	}

	if s[0] == 'I' {
		if (strLen > 2) && (s[1] == 'V') {
			Println("nextI equal to 'V'")
			sum += 4
			sum = romanToInt(s[2:strLen], sum)
		} else if (strLen > 2) && s[1] == 'X' {
			Println("nextI equal to 'X'")
			sum += 9
			sum = romanToInt(s[2:strLen], sum)
		} else if (strLen == 2) && (s[1] == 'V') {
			Println("nextI equal to 'V'")
			sum += 4
		} else if (strLen == 2) && s[1] == 'X' {
			Println("nextI equal to 'X'")
			sum += 9
		} else {
			Println("equal to 'I'")
			sum += 1
			sum = romanToInt(s[1:strLen], sum)
		}
	} else if s[0] == 'X' {
		if (strLen > 2) && (s[1] == 'L') {
			Println("nextI equal to 'L'")
			sum += 40
			sum = romanToInt(s[2:strLen], sum)
		} else if (strLen > 2) && s[1] == 'C' {
			Println("nextI equal to 'C'")
			sum += 90
			sum = romanToInt(s[2:strLen], sum)
		} else if (strLen == 2) && (s[1] == 'L') {
			Println("nextI equal to 'L'")
			sum += 40
		} else if (strLen == 2) && s[1] == 'C' {
			Println("nextI equal to 'C'")
			sum += 90
		} else {
			Println("equal to 'X'")
			sum += 10
			sum = romanToInt(s[1:strLen], sum)
		}
	} else if s[0] == 'C' {
		if (strLen > 2) && (s[1] == 'D') {
			Println("nextI equal to 'D'")
			sum += 400
			sum = romanToInt(s[2:strLen], sum)
		} else if (strLen > 2) && (s[1] == 'M') {
			Println("nextI equal to 'M'")
			sum += 900
			sum = romanToInt(s[2:strLen], sum)
		} else if (strLen == 2) && (s[1] == 'D') {
			Println("nextI equal to 'D'")
			sum += 400
		} else if (strLen == 2) && (s[1] == 'M') {
			Println("nextI equal to 'M'")
			sum += 900
		} else {
			Println("equal to 'C'")
			sum += 100
			sum = romanToInt(s[1:strLen], sum)
		}
	} else if s[0] == 'V' {
		Println("equal to 'V'")
		sum += 5
		sum = romanToInt(s[1:strLen], sum)
	} else if s[0] == 'L' {
		Println("equal to 'L'")
		sum += 50
		sum = romanToInt(s[1:strLen], sum)
	} else if s[0] == 'D' {
		Println("equal to 'D'")
		sum += 500
		sum = romanToInt(s[1:strLen], sum)
	} else if s[0] == 'M' {
		Println("equal to 'M'")
		sum += 1000
		sum = romanToInt(s[1:strLen], sum)
	}

	return sum
}
