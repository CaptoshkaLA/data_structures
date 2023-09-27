package main

import "fmt"

// MergeSort рекурсивно сортирует срез методом слияния
func MergeSort(arr []int) {
	n := len(arr)
	if n > 1 {
		mid := n / 2

		// Создаем два подмассива B и C
		B := make([]int, mid)
		copy(B, arr[:mid])
		C := make([]int, n-mid)
		copy(C, arr[mid:])

		// Рекурсивно сортируем подмассивы
		MergeSort(B)
		MergeSort(C)

		// Слияние отсортированных подмассивов B и C в массив A
		Merge(B, C, arr)
	}
}

// Merge объединяет два отсортированных подмассива B и C в массив A
func Merge(B, C, A []int) {
	i, j, k := 0, 0, 0
	p, q := len(B), len(C)

	for i < p && j < q {
		if B[i] <= C[j] {
			A[k] = B[i]
			i++
		} else {
			A[k] = C[j]
			j++
		}
		k++
	}

	// Копируем оставшиеся элементы из B или C
	for i < p {
		A[k] = B[i]
		i++
		k++
	}

	for j < q {
		A[k] = C[j]
		j++
		k++
	}
}

func TestCaseMergeSort() {
	arr := []int{64, 1, 25, 12, 46, 5, 22, 49, 11, 2}
	fmt.Println("Исходный массив:", arr)

	MergeSort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
