package main

import (
	"reflect"
	"testing"
	"github.com/Kostushka/go_algoritms/test"
)

type data struct {
	x        []int
	expected []int
}

func TestInsertionSort(t *testing.T) {
	// тестовые данные
	d := []data{
		{x: test.Test1, expected: test.Expected1},
		{x: test.Test2, expected: test.Expected2},
		{x: test.Test3, expected: test.Expected3},
	}

	for _, v := range d {
		// тестируемая функция
		insertionSort1(v.x)

		// проверка
		if reflect.DeepEqual(v.x, v.expected) == false {
			t.Errorf("expected %d, but got %d\n", v.expected, v.x)
		}
	}
}
