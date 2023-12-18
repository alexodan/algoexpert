package main

import (
	"reflect"
	"testing"
)

func Test_TwoNumberSum(t *testing.T) {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10

	if !reflect.DeepEqual(TwoNumberSum(array, targetSum), []int{-1, 11}) {
		t.Fatalf(`Error! %d`, TwoNumberSum(array, targetSum))
	}
}

func Test2_TwoNumberSum(t *testing.T) {
	array := []int{4, 6}
	targetSum := 10

	if !reflect.DeepEqual(TwoNumberSum(array, targetSum), []int{4, 6}) {
		t.Fatalf(`Error! %d`, TwoNumberSum(array, targetSum))
	}
}

func Test3_TwoNumberSum(t *testing.T) {
	array := []int{4, 6, 1}
	targetSum := 5

	if !reflect.DeepEqual(TwoNumberSum(array, targetSum), []int{1, 4}) {
		t.Fatalf(`Error! %d`, TwoNumberSum(array, targetSum))
	}
}
