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
	index := 0
	for i := len(negative); i > 0; i-- {
		negative[index] = i
		index++
	}
}

var sorted []int

func init() {
	sorted = make([]int, big)
	for i := 0; i < len(sorted); i++ {
		sorted[i] = i
	}
}

func BenchmarkQuickSortR(b *testing.B) {
	b.Run("random10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, random)
			quickSort(rcopy)
		} 
	})
	b.Run("random5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, random)
			quickSort(rcopy)
		}
	})	
	b.Run("random100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, random)
			quickSort(rcopy)
		}
	})
}

func BenchmarkLomutoSortR(b *testing.B) {
	b.Run("random10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, random)
			lomutoSort(rcopy)
		} 
	})
	b.Run("random5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, random)
			lomutoSort(rcopy)
		}
	})	
	b.Run("random100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, random)
			lomutoSort(rcopy)
		}
	})
}

func BenchmarkHoarSortR(b *testing.B) {
	b.Run("random10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, random)
			hoarSort(rcopy)
		} 
	})
	b.Run("random5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, random)
			hoarSort(rcopy)
		}
	})	
	b.Run("random100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, random)
			hoarSort(rcopy)
		}
	})
}

func BenchmarkQuickSortN(b *testing.B) {
	b.Run("negative10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, negative)
			quickSort(rcopy)
		} 
	})
	b.Run("negative5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, negative)
			quickSort(rcopy)
		}
	})	
	b.Run("negative100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, negative)
			quickSort(rcopy)
		}
	})
}

func BenchmarkLomutoSortN(b *testing.B) {
	b.Run("negative10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, negative)
			lomutoSort(rcopy)
		} 
	})
	b.Run("negative5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, negative)
			lomutoSort(rcopy)
		}
	})	
	b.Run("negative100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, negative)
			lomutoSort(rcopy)
		}
	})
}

func BenchmarkHoarSortN(b *testing.B) {
	b.Run("negative10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, negative)
			hoarSort(rcopy)
		} 
	})
	b.Run("negative5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, negative)
			hoarSort(rcopy)
		}
	})	
	b.Run("negative100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, negative)
			hoarSort(rcopy)
		}
	})
}

func BenchmarkQuickSortS(b *testing.B) {
	b.Run("sorted10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, sorted)
			quickSort(rcopy)
		} 
	})
	b.Run("sorted5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, sorted)
			quickSort(rcopy)
		}
	})	
	b.Run("sorted100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, sorted)
			quickSort(rcopy)
		}
	})
}

func BenchmarkLomutoSortS(b *testing.B) {
	b.Run("sorted10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, sorted)
			lomutoSort(rcopy)
		} 
	})
	b.Run("sorted5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, sorted)
			lomutoSort(rcopy)
		}
	})	
	b.Run("sorted100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, sorted)
			lomutoSort(rcopy)
		}
	})
}

func BenchmarkHoarSortS(b *testing.B) {
	b.Run("sorted10_000", func(b *testing.B) {
		rcopy := make([]int, big)
		for b.Loop() {
			copy(rcopy, sorted)
			hoarSort(rcopy)
		} 
	})
	b.Run("sorted5_000", func(b *testing.B) {
		rcopy := make([]int, middle)
		for b.Loop() {
			copy(rcopy, sorted)
			hoarSort(rcopy)
		}
	})	
	b.Run("sorted100", func(b *testing.B) {
		rcopy := make([]int, small)
		for b.Loop() {
			copy(rcopy, sorted)
			hoarSort(rcopy)
		}
	})
}
