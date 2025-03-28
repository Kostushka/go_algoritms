package main

import (
	"testing"
	"github.com/Kostushka/go_algoritms/test"
)

type data struct {
	x []int
	expected []int
}

func TestQuickSort(t *testing.T) {
	// тестовые данные
	d := []data{
		{x: test.Test1, expected: test.Expected1},	
		{x: test.Test2, expected: test.Expected2},	
		{x: test.Test3, expected: test.Expected3},	
	}
	
	for _, v := range d {
		// тестируемая функция
		quickSort(v.x)

		// проверка
		for i := 0; i < len(v.x); i++ {
			if v.x[i] != v.expected[i] {
				t.Errorf("expected %d, but got %d\n", v.expected, v.x)
				break
			}
		}
	}
}

func TestLomutoSort(t *testing.T) {
	// тестовые данные
	d := []data{
		{x: test.Test1, expected: test.Expected1},	
		{x: test.Test2, expected: test.Expected2},	
		{x: test.Test3, expected: test.Expected3},	
	}
	
	for _, v := range d {
		// тестируемая функция
		lomutoSort(v.x)

		// проверка
		for i := 0; i < len(v.x); i++ {
			if v.x[i] != v.expected[i] {
				t.Errorf("expected %d, but got %d\n", v.expected, v.x)
				break
			}
		}
	}
}

func TestHoarSort(t *testing.T) {
	// тестовые данные
	d := []data{
		{x: test.Test1, expected: test.Expected1},	
		{x: test.Test2, expected: test.Expected2},	
		{x: test.Test3, expected: test.Expected3},	
	}

	for _, v := range d {
		// тестируемая функция
		hoarSort(v.x)

		// проверка
		for i := 0; i < len(v.x); i++ {
			if v.x[i] != v.expected[i] {
				t.Errorf("expected %d, but got %d\n", v.expected, v.x)
				break
			}
		}
	}
}


