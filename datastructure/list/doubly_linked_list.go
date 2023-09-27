package list

import (
	"fmt"
)

// Структура для представления узла двусвязного списка
type NodeDLL struct {
	data interface{}
	next *NodeDLL
	prev *NodeDLL
}

// Структура для представления двусвязного списка
type DoublyLinkedList struct {
	head *NodeDLL
	tail *NodeDLL
	size int
}

// Функция для создания нового пустого списка
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Функция для добавления элемента в конец списка
func (ll *DoublyLinkedList) Append(data interface{}) {
	newNodeDLL := &NodeDLL{data: data}

	if ll.size == 0 {
		ll.head = newNodeDLL
		ll.tail = newNodeDLL
	} else {
		ll.tail.next = newNodeDLL
		newNodeDLL.prev = ll.tail
		ll.tail = newNodeDLL
	}

	ll.size++
}

// Функция для удаления элемента по индексу
func (ll *DoublyLinkedList) RemoveAtIndex(index int) {
	if index < 0 || index >= ll.size {
		return
	}

	currentNodeDLL := ll.head
	for i := 0; i < index; i++ {
		currentNodeDLL = currentNodeDLL.next
	}

	if currentNodeDLL.prev != nil {
		currentNodeDLL.prev.next = currentNodeDLL.next
	} else {
		ll.head = currentNodeDLL.next
	}

	if currentNodeDLL.next != nil {
		currentNodeDLL.next.prev = currentNodeDLL.prev
	} else {
		ll.tail = currentNodeDLL.prev
	}

	ll.size--
}

// Функция для получения элемента по индексу
func (ll *DoublyLinkedList) Get(index int) interface{} {
	if index < 0 || index >= ll.size {
		return nil
	}

	currentNodeDLL := ll.head
	for i := 0; i < index; i++ {
		currentNodeDLL = currentNodeDLL.next
	}

	return currentNodeDLL.data
}

// Функция для вставки элемента по индексу
func (ll *DoublyLinkedList) InsertAtIndex(index int, data interface{}) {
	if index < 0 || index > ll.size {
		return
	}

	newNodeDLL := &NodeDLL{data: data}

	if index == 0 {
		newNodeDLL.next = ll.head
		ll.head.prev = newNodeDLL
		ll.head = newNodeDLL
	} else if index == ll.size {
		ll.tail.next = newNodeDLL
		newNodeDLL.prev = ll.tail
		ll.tail = newNodeDLL
	} else {
		currentNodeDLL := ll.head
		for i := 0; i < index; i++ {
			currentNodeDLL = currentNodeDLL.next
		}

		newNodeDLL.prev = currentNodeDLL.prev
		newNodeDLL.next = currentNodeDLL
		currentNodeDLL.prev.next = newNodeDLL
		currentNodeDLL.prev = newNodeDLL
	}

	ll.size++
}

// Функция для печати списка
func (ll *DoublyLinkedList) PrintList() {
	currentNodeDLL := ll.head
	for currentNodeDLL != nil {
		fmt.Print(currentNodeDLL.data, " ")
		currentNodeDLL = currentNodeDLL.next
	}
	fmt.Println()
}

func TestCaseDoublyLinkedList() {
	ll := NewDoublyLinkedList()

	// Добавление элементов в конец списка
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	// Печать списка
	fmt.Println("Список после добавления элементов в конец:")
	ll.PrintList()

	// Получение элемента по индексу и печать
	index := 1
	element := ll.Get(index)
	fmt.Printf("Элемент на позиции %d: %v\n", index, element)

	// Вставка элемента по индексу и печать
	ll.InsertAtIndex(1, 4)
	fmt.Println("Список после вставки элемента 4 на позицию 1:")
	ll.PrintList()

	// Удаление элемента по индексу и печать
	ll.RemoveAtIndex(2)
	fmt.Println("Список после удаления элемента на позиции 2:")
	ll.PrintList()
}
