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
	"fmt"
)

func main() {

	fmt.Println("\nCircular Linked List\n")
	list2.TestCaseCircularLinkedList()

	fmt.Println("\nLinked List\n")
	list2.TestCaseLinkedList()

	fmt.Println("\nDoubly Linked List\n")
	list2.TestCaseDoublyLinkedList()

	fmt.Println("\nStack\n")
	stack.TestCaseStack()

	fmt.Println("\nQueue\n")
	queue.TestCaseQueue()

	fmt.Println("\nBinary Search Tree\n")
	tree2.TestCaseBinarySearchTree()

	fmt.Println("\nGraph\n")
	graph.TestCaseGraph()

	fmt.Println("\nHash Table\n")
	hash_table.TestCaseHashTable()

	fmt.Println("\nMerkle Tree\n")
	tree2.TestCaseMerkleTree()

	fmt.Println("\nSet\n")
	set.TestCaseSet()

	fmt.Println("\nBubble Sort\n")
	sorting.TestCaseBubbleSort()

	fmt.Println("\nSelection Sort\n")
	sorting.TestCaseSelectionSort()

	fmt.Println("\nInsertion Sort\n")
	sorting.TestCaseInsertionSort()

	fmt.Println("\nMerge Sort\n")
	sorting.TestCaseMergeSort()

	fmt.Println("\nHeap Sort\n")
	sorting.TestCaseHeapSort()
}
