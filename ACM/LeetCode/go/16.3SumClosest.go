package main

import (
	. "fmt"
	// "strings"
)

func main() {
	arr := []int{-2,4,-1,1}
	// arr := []int{-2,4}
	Println("arr: ", arr)
	target := 1
	Println("target: ", target)
	sum := closestSum(arr)
	Println("sum: ", sum)
}

func closestSum(arr []int) int {
	Println("arr: ", arr)
	if len(arr) < 3 {
		return 0
	}
	for i:=0;i<len(arr);i++{
		// Print("arr[i]: ",arr[i],"\t")
		for j:=i+1;j<len(arr);j++{
			// Print("arr[j]: ",arr[j],"\t")
			for k:=j+1;k<len(arr);k++{
				// Print("arr[k]: ",arr[k],"\t")
				Print("i,j,k: ",arr[i],arr[j],arr[k],"\t")
				Println()		
			}	
		}
		
	}
	Println()
	return 1
}
