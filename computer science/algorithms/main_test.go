package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var sortedList []int = []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	cases := map[int]int{
		1:  0,
		3:  1,
		4:  2,
		6:  3,
		7:  4,
		9:  5,
		10: 6,
		11: 7,
		13: 8,
		5:  -1,
	}
	for val, index := range cases {
		got := binarySearch(sortedList, val, 0, len(sortedList))
		if index != got {
			t.Errorf("For val : %d, Got index : %d, want : %d", val, got, index)
		}
	}
	lookingFor := 20
	index := binarySearch(sortedList, lookingFor, 0, len(sortedList)-1)
	if index != -1 {
		t.Errorf("index should be equal to -1")
	}
}

// const N = 50000

// func generateRandomInput() []int {
// 	var input []int
// 	for i := 0; i < N; i++ {
// 		input = append(input, rand.Int())
// 	}
// 	return input
// }

// var input = generateRandomInput()

// func BenchmarkMergeSort(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		mergeSort(input)
// 	}
// }
