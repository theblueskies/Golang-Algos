package bst

import (
	"fmt"
)

// Node stores the left and right child, the key and value of a given node.
type Node struct {
	key   int
	value string
	left  *Node
	right *Node
}

//BinarySearchTree stores the root and has method for operating on it.
type BinarySearchTree struct {
	root *Node
}

// Insert stores a key/value pair in a BinarySearchTree
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

//InOrderTraverse traverses a BinarySearchTree in order
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

// Search finds a key in the BinarySearchTree and returns the corresponding value
func (bst *BinarySearchTree) Search(key int) string {
	if bst.root != nil {
		result := search(bst.root, key)
		return result
	}
	return "Tree is empty"
}

func search(node *Node, key int) string {
	if node != nil {
		if key == node.key {
			return node.value
		} else if key < node.key {
			return search(node.left, key)
		} else if key > node.key {
			return search(node.right, key)
		}
	}
	return "Not found"
}
