package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []interface{}
}

// Создание нового стека
func NewStack() *Stack {
	return &Stack{}
}

// Помещение элемента на вершину стека
func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

// Удаление и возврат элемента с вершины стека
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("стек пуст")
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]
	return item, nil
}

// Возврат элемента с вершины стека без его удаления
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("стек пуст")
	}
	return s.data[len(s.data)-1], nil
}

// Проверка, пуст ли стек
func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Print() {
	for i := len(s.data) - 1; i >= 0; i-- {
		fmt.Println(s.data[i])
	}
}

func TestCaseStack() {
	fmt.Println("\nStack\n")

	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Содержимое стека:")
	stack.Print()

	// Проверка вершины стека
	top, err := stack.Peek()
	if err == nil {
		fmt.Println("Вершина стека:", top)
	} else {
		fmt.Println("Ошибка:", err)
	}

	// Извлечение элементов из стека
	for !stack.IsEmpty() {
		item, _ := stack.Pop()
		fmt.Println("Извлечено из стека:", item)
	}

	// Попытка извлечения из пустого стека
	_, err = stack.Pop()
	if err != nil {
		fmt.Println("Ошибка при извлечении из пустого стека:", err)
	}
}
