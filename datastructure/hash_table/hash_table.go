package hash_table

import (
	"fmt"
)

const tableSize = 10

type HashTable struct {
	data [tableSize][]interface{}
}

func NewHashTable() *HashTable {
	return &HashTable{}
}

func hash(key interface{}) int {
	// Простая хеш-функция методом деления
	hashCode := 0
	switch key.(type) {
	case int:
		hashCode = key.(int)
	case string:
		str := key.(string)
		for i := 0; i < len(str); i++ {
			hashCode += int(str[i])
		}
	default:
		// Для других типов можно реализовать свою хеш-функцию
		hashCode = int(key.(int))
	}
	return hashCode % tableSize
}

func (ht *HashTable) Insert(key interface{}, value interface{}) {
	index := hash(key)

	for ht.data[index] != nil {
		// Пробируем вперед, если слот занят
		index = (index + 1) % tableSize
	}

	// Вставляем значение в свободный слот
	ht.data[index] = []interface{}{key, value}
}

func (ht *HashTable) Get(key interface{}) interface{} {
	index := hash(key)

	for ht.data[index] != nil {
		if ht.data[index][0] == key {
			return ht.data[index][1]
		}
		// Продолжаем поиск вперед, если ключ не совпадает
		index = (index + 1) % tableSize
	}

	// Если ключ не найден, вернем nil
	return nil
}

func (ht *HashTable) Delete(key interface{}) {
	index := hash(key)

	for ht.data[index] != nil {
		if ht.data[index][0] == key {
			ht.data[index] = nil
			return
		}
		// Продолжаем поиск вперед, если ключ не совпадает
		index = (index + 1) % tableSize
	}
}

func TestCaseHashTable() {
	hashTable := NewHashTable()

	// Вставляем пары ключ-значение в хеш-таблицу
	hashTable.Insert("apple", 5)
	hashTable.Insert("banana", 7)
	hashTable.Insert("cherry", 9)

	// Получаем значения по ключам
	fmt.Println("apple:", hashTable.Get("apple"))   // Вывод: apple: 5
	fmt.Println("banana:", hashTable.Get("banana")) // Вывод: banana: 7
	fmt.Println("cherry:", hashTable.Get("cherry")) // Вывод: cherry: 9

	// Удаляем значение по ключу
	hashTable.Delete("banana")
	fmt.Println("banana:", hashTable.Get("banana")) // Вывод: banana: <nil>
}
