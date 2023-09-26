package main

import (
	"fmt"
)

// Определяем структуру элемента списка
type Node struct {
	Data int
	Next *Node
}

// Определяем структуру связного списка
type LinkedList struct {
	head *Node
}

// Добавление элемента в конец списка
func (ll *LinkedList) append(Data int) {
	newNode := &Node{Data: Data, Next: nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// Удаление элемента из списка
func (ll *LinkedList) delete(Data int) {
	if ll.head == nil {
		return
	}
	if ll.head.Data == Data {
		ll.head = ll.head.Next
		return
	}
	current := ll.head
	for current.Next != nil {
		if current.Next.Data == Data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// Получение элемента по индексу
func (ll *LinkedList) getByIndex(index int) (int, error) {
	current := ll.head
	for i := 0; i < index; i++ {
		if current == nil {
			return 0, fmt.Errorf("Индекс вне диапазона")
		}
		current = current.Next
	}
	if current == nil {
		return 0, fmt.Errorf("Индекс вне диапазона")
	}
	return current.Data, nil
}

// Вставка элемента по индексу
func (ll *LinkedList) insertAtIndex(index, Data int) error {
	newNode := &Node{Data: Data, Next: nil}
	if index == 0 {
		newNode.Next = ll.head
		ll.head = newNode
		return nil
	}
	current := ll.head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return fmt.Errorf("Индекс вне диапазона")
		}
		current = current.Next
	}
	if current == nil {
		return fmt.Errorf("Индекс вне диапазона")
	}
	newNode.Next = current.Next
	current.Next = newNode
	return nil
}

// Печать списка
func (ll *LinkedList) printList() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func TestCaseLinkedList() {
	list := LinkedList{}

	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)

	fmt.Println("Список после добавления элементов:")
	list.printList()

	list.delete(3)

	fmt.Println("Список после удаления элемента с значением 3:")
	list.printList()

	value, err := list.getByIndex(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Элемент с индексом 1: %d\n", value)
	}

	err = list.insertAtIndex(2, 5)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Список после вставки элемента 5 по индексу 2:")
	list.printList()
}
