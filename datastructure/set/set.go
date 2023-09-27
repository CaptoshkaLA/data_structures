package set

import (
	"fmt"
	"strings"
)

// Определение структуры Set (множество)
type Set struct {
	data map[string]bool
}

// Функция для создания нового множества
func NewSet() *Set {
	s := &Set{}
	s.data = make(map[string]bool)
	return s
}

// Функция для добавления элемента в множество
func (s *Set) Add(item string) {
	s.data[item] = true
}

// Функция для удаления элемента из множества
func (s *Set) Remove(item string) {
	delete(s.data, item)
}

// Функция для проверки наличия элемента в множестве
func (s *Set) Contains(item string) bool {
	_, exists := s.data[item]
	return exists
}

// Функция для получения размера множества
func (s *Set) Size() int {
	return len(s.data)
}

// Функция для печати содержимого множества
func (s *Set) PrintSet() {
	items := []string{}
	for item := range s.data {
		items = append(items, item)
	}
	fmt.Println("Set:", strings.Join(items, ", "))
}

func TestCaseSet() {
	// Создаем новое множество
	mySet := NewSet()

	// Добавляем элементы
	mySet.Add("apple")
	mySet.Add("banana")
	mySet.Add("cherry")

	mySet.PrintSet()

	// Проверяем наличие элементов
	fmt.Println("Contains 'apple':", mySet.Contains("apple")) // Должно быть true
	fmt.Println("Contains 'grape':", mySet.Contains("grape")) // Должно быть false

	// Получаем размер множества
	fmt.Println("Size of set:", mySet.Size()) // Должно быть 3

	// Удаляем элемент
	mySet.Remove("apple")

	mySet.PrintSet()
	fmt.Println("Size of set after removal:", mySet.Size()) // Должно быть 2
}
