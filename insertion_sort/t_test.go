package main

import (
	"reflect"
	"testing"
)

type data struct {
	x        []int
	expected []int
}

func TestInsertionSort(t *testing.T) {
	// тестовые данные
	d := []data{
		{x: []int{34, 2, 9, 1, 5, 12}, expected: []int{1, 2, 5, 9, 12, 34}},
		{x: []int{3, 64, 2, 6, 45, 4, 1, 67, 54, 86, 34, 24, 13}, expected: []int{1, 2, 3, 4, 6, 13, 24, 34, 45, 54, 64, 67, 86}},
		{x: []int{4, 4324, 2, 1, 54, 223, 23, 0}, expected: []int{0, 1, 2, 4, 23, 54, 223, 4324}},
	}

	for _, v := range d {
		// тестируемая функция
		insertionSort(v.x)

		// проверка
		if reflect.DeepEqual(v.x, v.expected) == false {
			t.Errorf("expected %d, but got %d\n", v.expected, v.x)
		}
	}
}
