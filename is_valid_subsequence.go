package main

func IsValidSubsequence(arr []int, sequence []int) bool {
	i := 0
	seqPointer := 0
	for i < len(arr) && seqPointer < len(sequence) {
		if arr[i] == sequence[seqPointer] {
			seqPointer++
			if seqPointer == len(sequence) {
				return true
			}
		}
		i++
	}
	return false
}
