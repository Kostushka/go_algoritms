package test

var Test1 = []int{34, 2, 9, 1, 5, 12, 9, 2, 5}
var Expected1 = []int{1, 2, 2, 5, 5, 9, 9, 12, 34}

var Test2 = []int{3, 64, 2, 6, 45, 4, 1, 67, 54, 86, 34, 64, 24, 13, 4}
var Expected2 = []int{1, 2, 3, 4, 4, 6, 13, 24, 34, 45, 54, 64, 64, 67, 86}

var Test3 = []int{4, 4324, 2, 1, 54, 0, 223, 23, 0}
var Expected3 = []int{0, 0, 1, 2, 4, 23, 54, 223, 4324}

func Equal(arr1, arr2 []int) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
