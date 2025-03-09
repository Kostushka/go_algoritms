package main

import (
	"testing"
	"math/rand"
)

const size = 100000

var random []int

func init() {
	rand.Seed(3)
	random = make([]int, size)
	for i := 0; i < len(random); i++ {
		random[i] = rand.Intn(700)
	}
}

var negative []int

func init() {
	negative = make([]int, size)
	j := 0
	for i := len(negative) - 1; i > 0; i-- {
		negative[j] = i
		j++
	}
}

var sorted []int

func init() {
	sorted = make([]int, size)
	for i := 0; i < len(sorted); i++ {
		sorted[i] = i
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	b.Run("random", func(b *testing.B) {
		for b.Loop() {
			selectionSort(random)
		}
	})
	b.Run("negative", func(b *testing.B) {
		for b.Loop() {
			selectionSort(negative)
		}
	})	
	b.Run("sorted", func(b *testing.B) {
		for b.Loop() {
			selectionSort(sorted)
		}
	})
}
