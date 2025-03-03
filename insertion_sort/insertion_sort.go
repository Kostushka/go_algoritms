package main

// import "fmt"

func main() {
	arr := []int{12, 3, 1, 5}
	// fmt.Println("before:", arr)
	insertionSort(arr)
	// fmt.Println("after:", arr)
}

// Вычислительная сложность: О(n^2)
// Пространственная сложность: O(1)
func insertionSort(arr []int) {
	var count int
	for i := 1; i < len(arr); i++ {
		swapped := true
		// предыдущий элемент
		prev := i - 1
		// текущий элемент
		curr := i

		count++
		// fmt.Println(count)
		// идем от текущего элемента до тех пор, пока есть предыдущий и пока текущий элемент меньше предыдущего
		for prev >= 0 && swapped {
			swapped = false
			// если текущий элемент меньше предыдущего, меняем их местами
			if arr[curr] < arr[prev] {
				arr[curr], arr[prev] = arr[prev], arr[curr]
				swapped = true
			}
			// fmt.Println(arr)
			prev--
			curr--
		}
	}
}
