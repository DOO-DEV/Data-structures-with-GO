package binarytree

import "fmt"

type Node struct {
	data  int
	right *Node
	left  *Node
}

type BinaryNode struct {
	root *Node
}

func (b *BinaryNode) Insert(item int) {
	if b.root == nil {
		b.root = &Node{data: item, left: nil, right: nil}
	} else {
		b.root.Insert(item)
	}
}

func (n *Node) Insert(item int) {
	if n == nil {
		return
	}
	if item <= n.data {
		if n.left == nil {
			n.left = &Node{data: item, left: nil, right: nil}
		} else {
			n.left.Insert(item)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: item, left: nil, right: nil}
		} else {
			n.right.Insert(item)
		}

	}
}

func (b *BinaryNode) PrintTree() {
	if b == nil {
		fmt.Println("Tree is Empty")
	}
	fmt.Printf("root are %d\n", b.root.data)
	left := b.root.left
	right := b.root.right
	fmt.Println(left, right)
	for left != nil && right != nil {
		fmt.Printf("left are %d\n", left.data)
		fmt.Printf("right are %d\n", right.data)
		left = left.left
		right = right.right
	}
}

func BinaryTree() {
	tree := &BinaryNode{}
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(4)

	tree.PrintTree()
}
