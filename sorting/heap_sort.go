package sorting

import (
	"fmt"
)

// Функция для построения кучи из исходного массива
func buildHeap(arr []int) {
	n := len(arr)

	// Начинаем с середины массива и идем к началу
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

// Функция для просеивания элемента вниз через кучу
func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	// Если левый потомок больше родителя
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// Если правый потомок больше родителя и левого потомка
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// Если largest не равен i, меняем элементы местами и рекурсивно просеиваем
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// Функция для сортировки массива с использованием пирамиды
func heapSort(arr []int) {
	n := len(arr)

	// Построение кучи (шаг 1)
	buildHeap(arr)

	// Выполнение сортировки (шаг 2)
	for i := n - 1; i > 0; i-- {
		// Перемещаем текущий максимальный элемент в конец
		arr[0], arr[i] = arr[i], arr[0]

		// Пересоздаем кучу для уменьшенного массива
		heapify(arr, i, 0)
	}
}

func TestCaseHeapSort() {
	fmt.Println("\nHeap Sort\n")

	arr := []int{64, 1, 25, 12, 46, 5, 22, 49, 11, 2}
	fmt.Println("Исходный массив:", arr)

	heapSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
