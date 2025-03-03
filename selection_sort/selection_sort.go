package main

// import "fmt"

func main() {
	arr := []int{21, 93, 23, 61, 82, 13, 12, 45}
	// fmt.Println("before:", arr)
	selectionSort(arr)
	// fmt.Println("after:", arr)
}

// Вычислительная сложность алгоритма: O(n^2)
// Пространственная сложность алгоритма: О(1)
func selectionSort(arr []int) {
	var min int
	var isChange bool

	// проходим по всем элементам кроме последнего
	for j := 0; j < len(arr)-1; j++ {
		min = j
		isChange = false
		// находим индекс самого маленького элемента из оставшихся относительно текущего
		for i := j + 1; i < len(arr); i++ {
			if arr[min] > arr[i] {
				min = i
				isChange = true
			}
		}
		// текущий элемент должен быть самым маленьким из оставшихся неотсортированных
		if isChange {
			arr[j], arr[min] = arr[min], arr[j]
		}
		// fmt.Println(arr)
	}
}
