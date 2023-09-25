package main

import (
	"fmt"
)

// Queue представляет собой структуру для хранения очереди
type Queue struct {
	data []interface{}
}

// Enqueue добавляет элемент в конец очереди
func (q *Queue) Enqueue(item interface{}) {
	q.data = append(q.data, item)
}

// Dequeue удаляет и возвращает элемент из начала очереди
func (q *Queue) Dequeue() interface{} {
	if len(q.data) == 0 {
		return nil // Очередь пуста
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

// IsEmpty проверяет, пуста ли очередь
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Size возвращает размер очереди
func (q *Queue) Size() int {
	return len(q.data)
}

// PrintQueue выводит все элементы в очереди
func (q *Queue) PrintQueue() {
	fmt.Print("Очередь: ")
	for _, item := range q.data {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}

func TestCaseQueue() {
	queue := Queue{}

	// Добавление элементов в очередь
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Вывод элементов в очереди
	queue.PrintQueue()

	// Проверка размера очереди
	fmt.Printf("Размер очереди: %d\n", queue.Size())

	// Вынимаем и добавляем элемент
	fmt.Printf("Вынимаем и добавляем элемент очереди: \n")
	queue.Dequeue()
	queue.Enqueue(4)

	// Вывод элементов в очереди
	queue.PrintQueue()

	// Извлечение элементов из очереди
	for !queue.IsEmpty() {
		item := queue.Dequeue()
		fmt.Printf("Извлечен элемент: %v\n", item)
	}

	// Проверка, что очередь пуста
	fmt.Printf("Очередь пуста: %v\n", queue.IsEmpty())
}
