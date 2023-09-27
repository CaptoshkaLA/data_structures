package sorting

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i

		// Находим индекс минимального элемента в оставшейся части массива
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		// Обмен значениями текущего элемента и минимального элемента
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func TestCaseSelectionSort() {
	fmt.Println("\nSelection Sort\n")

	arr := []int{64, 1, 25, 12, 46, 5, 22, 49, 11, 2}
	fmt.Println("Исходный массив:", arr)

	selectionSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
