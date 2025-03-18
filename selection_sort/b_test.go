package main

import (
	"testing"
	"math/rand"
)

const big = 10_000
const middle = 5_000
const small = 100

var random []int

func init() {
	rand.Seed(3)
	random = make([]int, big)
	for i := 0; i < len(random); i++ {
		random[i] = rand.Intn(700)
	}
}

var negative []int

func init() {
	negative = make([]int, big)
	j := 0
	for i := len(negative) - 1; i > 0; i-- {
		negative[j] = i
		j++
	}
}

var sorted []int

func init() {
	sorted = make([]int, big)
	for i := 0; i < len(sorted); i++ {
		sorted[i] = i
	}
}

func BenchmarkSelectionSortR(b *testing.B) {
	b.Run("random10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, random)
			selectionSort(rcopy)
		}
	})
	b.Run("random5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, random)
			selectionSort(rcopy)
		}
	})
	b.Run("random100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, random)
			selectionSort(rcopy)
		}
	})
}
func BenchmarkSelectionSortN(b *testing.B) {
	b.Run("negative10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, negative)
			selectionSort(rcopy)
		}
	})
	b.Run("negative5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, negative)
			selectionSort(rcopy)
		}
	})
	b.Run("negative100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, negative)
			selectionSort(rcopy)
		}
	})
}
func BenchmarkSelectionSortS(b *testing.B) {
	b.Run("sorted10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, sorted)
			selectionSort(rcopy)
		}
	})
	b.Run("sorted5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, sorted)
			selectionSort(rcopy)
		}
	})
	b.Run("sorted100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, sorted)
			selectionSort(rcopy)
		}
	})
}
