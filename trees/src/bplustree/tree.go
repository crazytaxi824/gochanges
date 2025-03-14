package bplustree

import (
	"fmt"
	"os"
	"slices"
)

const (
	// B+ 树的阶，非叶子节点的子节点个数范围是 [ceil(order/2), order]
	// (order-1) keys, (order) children.
	order = 4 // 最多3个key, 4个children, 相当于 MaxKey=3
)

type BPlusTree struct {
	Root *Node
}

func NewBPlusTree() *BPlusTree {
	// Create a new leaf node as the root
	root := NewNode(true)
	return &BPlusTree{Root: root}
}

// findLeafNode finds the leaf node that would contain the given key
// for insert node or search node
func (t *BPlusTree) findLeafNode(key int) *Node {
	// Start from the root and traverse down to the leaf node
	node := t.Root

	// Traverse down the tree until we reach a leaf node
	for !node.IsLeaf {
		// Find the right child to follow
		i := slices.IndexFunc(node.Keys, func(k int) bool {
			return key < k
		})
		if i < 0 {
			i = len(node.Keys)
		}

		// Follow the child pointer
		node = node.Children[i]
	}

	return node
}

func (t *BPlusTree) Search(key int) (*Node, error) {
	leaf := t.findLeafNode(key)

	// Now we are at a leaf node, search for the key
	if slices.Contains(leaf.Keys, key) {
		return leaf, nil // Key found
	}

	// Key not found
	return nil, os.ErrNotExist
}

func (t *BPlusTree) Insert(key int) error {
	// Find the leaf node where the key should be inserted
	leaf := t.findLeafNode(key)

	// Check if the key already exists
	if slices.Contains(leaf.Keys, key) {
		return os.ErrExist
	}

	// insert key
	leaf.Keys = append(leaf.Keys, key)
	slices.Sort(leaf.Keys)

	// Handle the case where the leaf node is full
	if len(leaf.Keys) >= order {
		newNode, pk := leaf.SplitNode()
		newRoot := leaf.InsertIntoParent(newNode, pk)
		if newRoot != nil {
			t.Root = newRoot
		}
	}

	return nil
}

// PrintTree prints the tree structure for debugging
func PrintTree(node *Node, level int) {
	if node == nil {
		return
	}

	indent := ""
	for range level {
		indent += "\t"
	}

	fmt.Printf("%sNode(", indent)
	if node.IsLeaf {
		fmt.Printf("Leaf): Keys: %v", node.Keys)
		if node.Parent != nil {
			fmt.Printf(", parent: %v", node.Parent.Keys)
		}
		if node.Next != nil {
			fmt.Printf(", next: %v", node.Next.Keys)
		}
		fmt.Println()
	} else {
		fmt.Printf("Internal): Keys: %v", node.Keys)
		if node.Parent != nil {
			fmt.Printf(", parent: %v", node.Parent.Keys)
		}
		fmt.Println()
		for _, child := range node.Children {
			PrintTree(child, level+1)
		}
	}
}
