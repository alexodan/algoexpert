package main

import (
	"reflect"
	"testing"
)

func TestIsValidSubsequence(t *testing.T) {
	testCases := []struct {
		array    []int
		sequence []int
		expected bool
	}{
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, -1, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 25, 6, -1, 8, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 6, -1, 8, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{22, 25, 6}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{1, 6, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, -1, 8, 10}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{25}, true},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1}, true},
		{[]int{5, 1, 22, 25, 6, -1, 8, 10}, []int{5, 1, 22, 25, 6, -1, 8, 10, 12}, false},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result := IsValidSubsequence(tc.array, tc.sequence)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("For array %v and sequence %v, expected %v, got %v", tc.array, tc.sequence, tc.expected, result)
			}
		})
	}
}
