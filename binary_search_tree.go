package main

import (
	"fmt"
)

// Узел бинарного дерева поиска
type NodeBST struct {
	Data  interface{}
	Left  *NodeBST
	Right *NodeBST
}

// Бинарное дерево поиска
type BST struct {
	Root *NodeBST
}

// Вставка элемента в BST
func (bst *BST) Insert(data interface{}) {
	bst.Root = insertRec(bst.Root, data)
}

func insertRec(root *NodeBST, data interface{}) *NodeBST {
	if root == nil {
		return &NodeBST{Data: data}
	}

	if data.(int) < root.Data.(int) {
		root.Left = insertRec(root.Left, data)
	} else if data.(int) > root.Data.(int) {
		root.Right = insertRec(root.Right, data)
	}

	return root
}

// Поиск элемента в BST
func (bst *BST) Search(data interface{}) bool {
	return searchRec(bst.Root, data)
}

func searchRec(root *NodeBST, data interface{}) bool {
	if root == nil {
		return false
	}

	if data.(int) < root.Data.(int) {
		return searchRec(root.Left, data)
	} else if data.(int) > root.Data.(int) {
		return searchRec(root.Right, data)
	} else {
		return true
	}
}

// Удаление элемента из BST
func (bst *BST) Delete(data interface{}) {
	bst.Root = deleteRec(bst.Root, data)
}

func deleteRec(root *NodeBST, data interface{}) *NodeBST {
	if root == nil {
		return root
	}

	if data.(int) < root.Data.(int) {
		root.Left = deleteRec(root.Left, data)
	} else if data.(int) > root.Data.(int) {
		root.Right = deleteRec(root.Right, data)
	} else {
		// Узел с одним или без детей
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Узел с двумя детьми: найти наименьший элемент в правом поддереве
		root.Data = minValue(root.Right)

		// Удалить наименьший элемент в правом поддереве
		root.Right = deleteRec(root.Right, root.Data)
	}

	return root
}

func minValue(node *NodeBST) interface{} {
	minValue := node.Data
	for node.Left != nil {
		minValue = node.Left.Data
		node = node.Left
	}
	return minValue
}

// Печать BST в виде дерева (инфиксный обход)
func (bst *BST) PrintTree() {
	printTreeRec(bst.Root, 0)
}

func printTreeRec(node *NodeBST, depth int) {
	if node == nil {
		return
	}

	// Рекурсивно обойти правое поддерево
	printTreeRec(node.Right, depth+1)

	// Вывести отступ для текущего уровня глубины
	for i := 0; i < depth; i++ {
		fmt.Print("\t")
	}

	// Вывести данные узла
	fmt.Println(node.Data)

	// Рекурсивно обойти левое поддерево
	printTreeRec(node.Left, depth+1)
}

func TestCaseBinarySearchTree() {
	bst := BST{}
	data := []int{50, 30, 70, 20, 40, 60, 80}

	for _, value := range data {
		bst.Insert(value)
	}

	fmt.Println("Дерево:")
	bst.PrintTree()

	searchKey := 30
	if bst.Search(searchKey) {
		fmt.Printf("Элемент %d найден в дереве.\n", searchKey)
	} else {
		fmt.Printf("Элемент %d не найден в дереве.\n", searchKey)
	}

	deleteKey := 40
	bst.Delete(deleteKey)
	fmt.Printf("Удален элемент %d из дерева.\n", deleteKey)

	fmt.Println("Дерево после удаления:")
	bst.PrintTree()
}
