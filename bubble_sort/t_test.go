package main

import (
	"testing"
	"github.com/Kostushka/go_algoritms/test"
)

type data struct {
	x        []int
	expected []int
}

func TestBubbleSort(t *testing.T) {
	// тестовые данные
	d := []data{
		{x: test.Test1, expected: test.Expected1},
		{x: test.Test2, expected: test.Expected2},
		{x: test.Test3, expected: test.Expected3},
	}

	for _, v := range d {
		// тестируемая функция
		bubbleSort(v.x)

		// проверка
		if !test.Equal(v.x, v.expected) {
			t.Errorf("expected %d, but got %d\n", v.expected, v.x)
		}
	}
}
