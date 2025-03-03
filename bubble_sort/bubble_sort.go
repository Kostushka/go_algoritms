package main

// import "fmt"

func main() {
	arr := []int{21, 93, 23, 61, 82, 13, 12, 45}
	// fmt.Println("before: ", arr)
	bubbleSort(arr)
	// fmt.Println("after: ", arr)
}

// Вычислительная сложность алгоритма: O(n^2)
// Пространственная сложность алгоритма: О(1)
func bubbleSort(arr []int) {
	var count int
	var swapped bool = true
	for i := 0; i < len(arr) - 1 && swapped; i++ {
		count++
		// fmt.Println(count)
		swapped = false
		for j := 0; j < len(arr) - 1; j++ {
			if arr[j] > arr[j+1] {
				swapped = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			// fmt.Println(arr)
		}
	}
}
