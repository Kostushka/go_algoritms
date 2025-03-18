package main

// import "fmt"

func main() {
	arr := []int{34, 2, 9, 1, 5, 12}
	// fmt.Println("before:", arr)
	insertionSort1(arr)
	// fmt.Println("after:", arr)
}

// Вычислительная сложность: О(n^2)
// Пространственная сложность: O(1)
func insertionSort1(arr []int) {
	// var count int
	for i := 1; i < len(arr); i++ {
		// предыдущий элемент
		prev := i - 1
		// текущий элемент
		curr := i

		// элемент, который надо поставить на место
		sortable := arr[i]

		// count++
		// fmt.Println(count)
		// идем от текущего элемента до тех пор, пока есть предыдущий
		for prev >= 0 {
			// если элемент, который надо поставить на место, меньше предыдущего, сдвигаем предыдущий на одну позицию вперед
			if sortable < arr[prev] {
				arr[curr] = arr[prev]
				prev--
				curr--
				continue
			}
			// иначе ставим элемент на место
			arr[curr] = sortable
			break
		}
		// если дошли до начала массива, значит элемент, который надо поставить на место, должен быть первым
		if prev < 0 {
			arr[0] = sortable
		}
	}
}

// Вычислительная сложность: О(n^2)
// Пространственная сложность: O(1)
func insertionSort2(arr []int) {
	// var count int
external:
	for i := 1; i < len(arr); i++ {
		// предыдущий элемент
		prev := i - 1
		// текущий элемент
		curr := i

		// элемент, который надо поставить на место
		sortable := arr[i]

		// count++
		// fmt.Println(count)
		// идем от текущего элемента до тех пор, пока есть предыдущий
		for prev >= 0 {
			// если элемент, который надо поставить на место, меньше предыдущего, сдвигаем предыдущий на одну позицию вперед
			if sortable < arr[prev] {
				arr[curr] = arr[prev]
				prev--
				curr--
				continue
			}
			// иначе ставим элемент на место
			arr[curr] = sortable
			continue external
		}
		// если дошли до начала массива, значит элемент, который надо поставить на место, должен быть первым
		arr[0] = sortable
	}
}
