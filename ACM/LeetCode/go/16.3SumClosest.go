package main

import (
	. "fmt"
)

func main() {
	arr := []int{-3,-2,-5,3,-4}
	target := -1
	sum := closestSum(arr, target)
	Println("sum: ", sum)
}

func closestSum(arr []int, target int) int {
	if len(arr) < 3 {
		return 0
	}
	flag := arr[0] + arr[1] + arr[2]
	gap := toNonNegative(flag - target)
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			for k:=j+1;k<len(arr);k++{
				realtmp := arr[i] + arr[j] + arr[k]
				tmp := toNonNegative(realtmp - target)
				if (tmp < gap) {
					gap = tmp
					flag = realtmp
				}
			}	
		}		
	}
	return flag
}

func toNonNegative(x int) int {
	if (x < 0) {
		return -x
	}
	return x
}