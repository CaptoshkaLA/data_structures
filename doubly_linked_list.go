package main

import (
	"fmt"
)

// Структура для представления узла двусвязного списка
type NodeD struct {
	data interface{}
	next *NodeD
	prev *NodeD
}

// Структура для представления двусвязного списка
type DoublyLinkedList struct {
	head *NodeD
	tail *NodeD
	size int
}

// Функция для создания нового пустого списка
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Функция для добавления элемента в конец списка
func (ll *DoublyLinkedList) Append(data interface{}) {
	newNodeD := &NodeD{data: data}

	if ll.size == 0 {
		ll.head = newNodeD
		ll.tail = newNodeD
	} else {
		ll.tail.next = newNodeD
		newNodeD.prev = ll.tail
		ll.tail = newNodeD
	}

	ll.size++
}

// Функция для удаления элемента по индексу
func (ll *DoublyLinkedList) RemoveAtIndex(index int) {
	if index < 0 || index >= ll.size {
		return
	}

	currentNodeD := ll.head
	for i := 0; i < index; i++ {
		currentNodeD = currentNodeD.next
	}

	if currentNodeD.prev != nil {
		currentNodeD.prev.next = currentNodeD.next
	} else {
		ll.head = currentNodeD.next
	}

	if currentNodeD.next != nil {
		currentNodeD.next.prev = currentNodeD.prev
	} else {
		ll.tail = currentNodeD.prev
	}

	ll.size--
}

// Функция для получения элемента по индексу
func (ll *DoublyLinkedList) Get(index int) interface{} {
	if index < 0 || index >= ll.size {
		return nil
	}

	currentNodeD := ll.head
	for i := 0; i < index; i++ {
		currentNodeD = currentNodeD.next
	}

	return currentNodeD.data
}

// Функция для вставки элемента по индексу
func (ll *DoublyLinkedList) InsertAtIndex(index int, data interface{}) {
	if index < 0 || index > ll.size {
		return
	}

	newNodeD := &NodeD{data: data}

	if index == 0 {
		newNodeD.next = ll.head
		ll.head.prev = newNodeD
		ll.head = newNodeD
	} else if index == ll.size {
		ll.tail.next = newNodeD
		newNodeD.prev = ll.tail
		ll.tail = newNodeD
	} else {
		currentNodeD := ll.head
		for i := 0; i < index; i++ {
			currentNodeD = currentNodeD.next
		}

		newNodeD.prev = currentNodeD.prev
		newNodeD.next = currentNodeD
		currentNodeD.prev.next = newNodeD
		currentNodeD.prev = newNodeD
	}

	ll.size++
}

// Функция для печати списка
func (ll *DoublyLinkedList) PrintList() {
	currentNodeD := ll.head
	for currentNodeD != nil {
		fmt.Print(currentNodeD.data, " ")
		currentNodeD = currentNodeD.next
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
