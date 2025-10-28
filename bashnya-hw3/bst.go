package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	item  int
	left  *Node
	right *Node
}

func (t *Tree) Insert(value int) {
	t.root = insert(t.root, value)
}

func insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{item: value}
	}
	if value < node.item {
		node.left = insert(node.left, value)
	} else if value > node.item {
		node.right = insert(node.right, value)
	}
	return node
}

func (t *Tree) Remove(value int) {
	t.root = remove(t.root, value)
}

func remove(node *Node, value int) *Node {
	if node == nil {
		return nil
	}
	if value < node.item {
		node.left = remove(node.left, value)
	} else if value > node.item {
		node.right = remove(node.right, value)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

	}
}

func (t *Tree) Find(value int) bool {
	return find(t.root, value)
}

func find(node *Node, value int) bool {
	if node == nil {
		return false
	}
	if value == node.item {
		return true
	} else if value < node.item {
		return find(node.left, value)
	} else {
		return find(node.right, value)
	}
}

func main() {
	tree := Tree{nil}
	tree.Insert(5)
	tree.Insert(10)
	fmt.Println(tree.Search(0))
}
