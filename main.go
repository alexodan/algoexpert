package main

import (
	"fmt"
)

func main() {
	// Two number sum
	// r1 := TwoNumberSum([]int{3, 5, -4, 8, 11, 1, -1, 6}, 10)
	// fmt.Println(r1)

	arr := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}

	r2 := IsValidSubsequence(arr, sequence)
	fmt.Println(r2)
}
