package sorting

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		j := i - 1
		for j >= 0 && arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j]
			j--
		}
	}
}

func TestCaseInsertionSort() {
	arr := []int{64, 1, 25, 12, 46, 5, 22, 49, 11, 2}
	fmt.Println("Неотсортированный массив:", arr)
	insertionSort(arr)
	fmt.Println("Отсортированный массив:", arr)
}
