package main

import (
	"crypto/sha256"
	"fmt"
)

// TreeNode представляет узел дерева Меркла.
type TreeNode struct {
	Hash  string
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode создает новый узел с заданными данными и хеширует их.
func NewTreeNode(data string) *TreeNode {
	hash := calculateHash(data)
	return &TreeNode{Hash: hash}
}

// calculateHash вычисляет хеш для данных.
func calculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// buildMerkleTree строит дерево Меркла из списка данных.
func buildMerkleTree(data []string) *TreeNode {
	var nodes []*TreeNode

	// Создаем листовые узлы и добавляем их в список.
	for _, d := range data {
		nodes = append(nodes, NewTreeNode(d))
	}

	// Строим дерево Меркла путем комбинирования узлов в пары и создания родительских узлов.
	for len(nodes) > 1 {
		var newNodes []*TreeNode
		for i := 0; i < len(nodes); i += 2 {
			node1 := nodes[i]
			var node2 *TreeNode
			if i+1 < len(nodes) {
				node2 = nodes[i+1]
			}
			parent := NewTreeNode(node1.Hash + node2.Hash)
			parent.Left = node1
			parent.Right = node2
			newNodes = append(newNodes, parent)
		}
		nodes = newNodes
	}

	// Возвращаем корневой узел дерева Меркла.
	return nodes[0]
}

// printMerkleTree выводит дерево Меркла в более наглядной форме.
func printMerkleTree(node *TreeNode, indent string, isLeft bool) {
	if node == nil {
		return
	}

	// Выбираем символ для отображения узла в зависимости от его положения.
	var nodeSymbol string
	if isLeft {
		nodeSymbol = "├──"
	} else {
		nodeSymbol = "└──"
	}

	// Выводим узел с хешем.
	fmt.Println(indent + nodeSymbol + " " + node.Hash)

	// Рекурсивно выводим левое и правое поддерево.
	printMerkleTree(node.Left, indent+"│   ", true)
	printMerkleTree(node.Right, indent+"│   ", false)
}

func TestCaseMerkleTree() {
	data := []string{"block1", "block2", "block3", "block4"}

	// Строим дерево Меркла из данных.
	root := buildMerkleTree(data)

	// Выводим дерево Меркла.
	fmt.Println("Merkle Tree:")
	printMerkleTree(root, "", false)

	// Пример удаления узла (заменяем "block2" на "block5").
	replaceNode(root.Left.Left, "block5")

	// Выводим обновленное дерево Меркла.
	fmt.Println("\nUpdated Merkle Tree:")
	printMerkleTree(root, "", false)
}

// replaceNode заменяет данные в узле и обновляет хеши вверх по дереву.
func replaceNode(node *TreeNode, newData string) {
	if node == nil {
		return
	}

	node.Hash = calculateHash(newData)

	// Рекурсивно обновляем хеши вверх по дереву.
	if node.Left != nil {
		replaceNode(node.Left, newData)
	}
	if node.Right != nil {
		replaceNode(node.Right, newData)
	}
}
