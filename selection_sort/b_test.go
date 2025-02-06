package main

import "testing"

func BenchmarkSelectionSort(b *testing.B) {
	x := []int{3, 64, 2, 6, 45, 4, 1, 67, 54, 86, 34, 24, 13}

	for i := 0; i < b.N; i++ {
		selectionSort(x)
	}
}
