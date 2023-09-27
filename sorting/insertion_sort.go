package sorting

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	// Проходим по массиву, начиная с элемента под индексом 1
	for i := 1; i < n; i++ {
		j := i - 1
		// Перемещаем элементы, пока не встретим место для вставки текущего элемента
		for j >= 0 && arr[j] > arr[j+1] {
			arr[j], arr[j+1] = arr[j+1], arr[j] // Меняем местами элементы
			j--
		}
	}
}

func TestCaseInsertionSort() {
	fmt.Println("\nInsertion Sort\n")

	arr := []int{64, 1, 25, 12, 46, 5, 22, 49, 11, 2}
	fmt.Println("Неотсортированный массив:", arr)

	insertionSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
