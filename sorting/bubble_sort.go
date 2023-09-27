package sorting

import "fmt"

func BubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		// Внутренний цикл - перебор элементов массива
		// (каждая итерация перемещает максимальный элемент в конец)
		for j := 0; j < n-1-i; j++ {
			// Сравниваем текущий элемент со следующим
			if arr[j+1] < arr[j] {
				// Если следующий элемент меньше текущего, меняем их местами
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func TestCaseBubbleSort() {
	fmt.Println("\nBubble Sort\n")

	arr := []int{64, 1, 25, 12, 46, 5, 22, 49, 11, 2}
	fmt.Println("Исходный массив:", arr)

	BubbleSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
