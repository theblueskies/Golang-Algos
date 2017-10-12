package main

import (
	"fmt"
)

type Node struct {
	key   int
	value string
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(key int, value string) {
	newNode := &Node{key: key, value: value}
	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
}

func insertNode(currentNode, newNode *Node) {
	if newNode.key < currentNode.key {
		if currentNode.left == nil {
			currentNode.left = newNode
		} else {
			insertNode(currentNode.left, newNode)
		}
	} else {
		if currentNode.right == nil {
			currentNode.right = newNode
		} else {
			insertNode(currentNode.right, newNode)
		}
	}
}

func (bst *BinarySearchTree) InOrderTraverse() {
	inOrderTraverse(bst.root)
}

func inOrderTraverse(node *Node) {
	if node != nil {

		inOrderTraverse(node.left)
		fmt.Println(node.value)
		inOrderTraverse(node.right)
	}
}

func main() {
	bst := &BinarySearchTree{}

	bst.Insert(8, "8")
	bst.Insert(5, "5")
	bst.Insert(3, "3")
	bst.Insert(7, "7")
	bst.Insert(10, "10")
	bst.Insert(11, "11")
	bst.Insert(9, "9")

	bst.InOrderTraverse()
}
