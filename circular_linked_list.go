package main

import (
	"fmt"
)

// Структура для элемента списка
type Node struct {
	Data int
	Next *Node
}

// Структура для кольцевого списка
type CircularLinkedList struct {
	Head *Node
	Tail *Node
}

// Метод добавления элемента в конец списка
func (cll *CircularLinkedList) Add(data int) {
	newNode := &Node{Data: data}
	if cll.Head == nil {
		cll.Head = newNode
		cll.Tail = newNode
	} else {
		cll.Tail.Next = newNode
		cll.Tail = newNode
	}
	// Связываем последний элемент с головой, чтобы сделать список кольцевым
	cll.Tail.Next = cll.Head
}

// Метод удаления элемента из списка по значению
func (cll *CircularLinkedList) Remove(data int) {
	if cll.Head == nil {
		return
	}

	// Если элемент находится в голове, то обновляем голову
	if cll.Head.Data == data {
		if cll.Head == cll.Tail {
			cll.Head = nil
			cll.Tail = nil
		} else {
			cll.Tail.Next = cll.Head.Next
			cll.Head = cll.Head.Next
		}
		return
	}

	// Иначе ищем элемент и удаляем его
	current := cll.Head
	for current.Next != nil && current.Next != cll.Head {
		if current.Next.Data == data {
			if current.Next == cll.Tail {
				cll.Tail = current
			}
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

// Метод вывода всех элементов списка
func (cll *CircularLinkedList) Print() {
	if cll.Head == nil {
		return
	}

	current := cll.Head
	for {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
		if current == cll.Head {
			break
		}
	}
	fmt.Println()
}

// Метод получения элемента по индексу
func (cll *CircularLinkedList) GetByIndex(index int) *Node {
	if cll.Head == nil || index < 0 {
		return nil
	}

	current := cll.Head
	for i := 0; i < index; i++ {
		current = current.Next
		if current == cll.Head {
			return nil
		}
	}
	return current
}

// Метод вставки элемента по индексу
func (cll *CircularLinkedList) InsertByIndex(index, data int) {
	if index == 0 {
		newNode := &Node{Data: data, Next: cll.Head}
		cll.Tail.Next = newNode
		cll.Head = newNode
		return
	}

	prevNode := cll.GetByIndex(index - 1)
	if prevNode == nil {
		fmt.Println("Невозможно вставить элемент. Неверный индекс.")
		return
	}

	newNode := &Node{Data: data, Next: prevNode.Next}
	prevNode.Next = newNode
	if prevNode == cll.Tail {
		cll.Tail = newNode
	}
}

// Метод проверки, что первый элемент указывает на последний, а последний на первый
func (cll *CircularLinkedList) IsCircular() bool {
	if cll.Head == nil {
		return false
	}
	return cll.Tail.Next == cll.Head
}

func TestCaseCircularLinkedList() {
	circularList := &CircularLinkedList{}

	circularList.Add(1)
	circularList.Add(2)
	circularList.Add(3)
	circularList.Add(4)

	fmt.Println("Список:")
	circularList.Print()

	// Удаление элемента из списка
	circularList.Remove(2)
	fmt.Println("После удаления элемента 2:")
	circularList.Print()

	// Вставка элемента по индексу
	index := 1
	dataToInsert := 5
	circularList.InsertByIndex(index, dataToInsert)
	fmt.Printf("После вставки элемента %d по индексу %d:\n", dataToInsert, index)
	circularList.Print()

	// Получение элемента по индексу
	index = 2
	element := circularList.GetByIndex(index)
	if element != nil {
		fmt.Printf("Элемент с индексом %d: %d\n", index, element.Data)
	} else {
		fmt.Printf("Элемент с индексом %d не найден\n", index)
	}

	// Вывод головы списка и следующего элемента по индексу
	fmt.Printf("Голова списка: %d\n", circularList.Head.Data)
	nextElement := circularList.GetByIndex(1)
	if nextElement != nil {
		fmt.Printf("Следующий элемент после индекса 0: %d\n", nextElement.Data)
	}

	// Проверка, что список является кольцевым
	fmt.Printf("Список кольцевой: %v\n", circularList.IsCircular())
}
