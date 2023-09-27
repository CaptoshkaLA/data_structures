package main

import "fmt"

// Функция для сортировки пузырьком
func BubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j+1] < arr[j] {
				// Обмен значениями
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func TestCaseBubbleSort() {
	// Пример использования сортировки пузырьком
	arr := []int{64, 1, 25, 12, 46, 5, 22, 49, 11, 2}
	fmt.Println("Исходный массив:", arr)

	// Вызов функции сортировки
	BubbleSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
