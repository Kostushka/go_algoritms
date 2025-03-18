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

	// проходим по всем элементам кроме последнего
	loop:
	for i := 0; i < len(arr)-1; i++ {
		min = i
		// находим индекс самого маленького элемента из оставшихся относительно текущего
		for j := i + 1; j < len(arr); j++ {
			if min > 0 && (arr[min] == arr[min-1] || arr[min] == arr[min-1]+1) {
				continue loop
			}
			if arr[min] > arr[j] {
				min = j
			}
		}
		// текущий элемент должен быть самым маленьким из оставшихся неотсортированных
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
		// fmt.Println(arr)
	}
}
