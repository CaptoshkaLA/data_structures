package main

import (
	"Sandbox_/datastructure/graph"
	"Sandbox_/datastructure/hash_table"
	list2 "Sandbox_/datastructure/list"
	"Sandbox_/datastructure/queue"
	"Sandbox_/datastructure/set"
	"Sandbox_/datastructure/stack"
	tree2 "Sandbox_/datastructure/tree"
	"Sandbox_/sorting"
)

func main() {

	list2.TestCaseCircularLinkedList()

	list2.TestCaseLinkedList()

	list2.TestCaseDoublyLinkedList()

	stack.TestCaseStack()

	queue.TestCaseQueue()

	tree2.TestCaseBinarySearchTree()

	graph.TestCaseGraph()

	hash_table.TestCaseHashTable()

	tree2.TestCaseMerkleTree()

	set.TestCaseSet()

	sorting.TestCaseBubbleSort()

	sorting.TestCaseSelectionSort()

	sorting.TestCaseInsertionSort()

	sorting.TestCaseMergeSort()

	sorting.TestCaseHeapSort()
}
