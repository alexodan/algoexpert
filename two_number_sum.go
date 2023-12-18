package main

import (
	"sort"
)

func TwoNumberSum(array []int, target int) []int {
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	i := 0
	j := len(array) - 1
	for i < j {
		sum := array[i] + array[j]
		if sum == target {
			return []int{array[i], array[j]}
		}
		if sum < target {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
