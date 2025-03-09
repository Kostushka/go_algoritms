package test

func Equal(arr1, arr2 []int) bool {
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
