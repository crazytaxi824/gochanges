//        3   5
//     /    |    \
//   1 2   3 4   5 6 7

package bplustree

import (
	"slices"
)

// Node represents a node in the B+Tree
type Node struct {
	IsLeaf   bool
	Keys     []int
	Children []*Node // Only used for internal nodes
	Next     *Node   // Only used for leaf nodes (for range queries)
	Parent   *Node   // Parent reference
}

// NewNode creates a new node
func NewNode(isLeaf bool) *Node {
	return &Node{
		IsLeaf:   isLeaf,
		Keys:     make([]int, 0, order), // 多一个位置为了 split
		Children: make([]*Node, 0, order+1),
		Next:     nil,
		Parent:   nil,
	}
}

// SplitLeafNode splits a leaf node that has reached maximum capacity
func SplitLeafNode(node *Node) (newRightNode *Node, promotedKey int) {
	if !node.IsLeaf {
		panic("Can't use SplitLeafNode on an internal node")
	}

	// Create a new leaf node
	// NOTE: 这里不设置 Parent, 因为 node 可能是 root 节点, 没有 Parent. Parent 设置放在后面.
	rightNode := NewNode(true)

	// Calculate split point - middle of the node
	splitIndex := len(node.Keys) / 2

	// Move half of the keys and values to the new node
	rightNode.Keys = append(rightNode.Keys, node.Keys[splitIndex:]...)

	// Update the original node's keys and values
	node.Keys = node.Keys[:splitIndex]

	// Handle the linked list of leaf nodes for range queries
	rightNode.Next = node.Next
	node.Next = rightNode

	// Get the key to promote to the parent
	promotedKey = rightNode.Keys[0]

	return rightNode, promotedKey
}

// SplitInternalNode splits an internal node that has reached maximum capacity
func SplitInternalNode(node *Node) (newRightNode *Node, promotedKey int) {
	if node.IsLeaf {
		panic("Can't use SplitInternalNode on a leaf node")
	}

	// Create a new internal node
	// NOTE: 这里不设置 Parent, 因为 node 可能是 root 节点, 没有 Parent. Parent 设置放在后面.
	rightNode := NewNode(false)

	// Calculate split point - middle of the node
	splitIndex := len(node.Keys) / 2

	// Get the key to promote to the parent
	promotedKey = node.Keys[splitIndex]

	// Move keys after the middle to the new node (excluding the middle key)
	rightNode.Keys = append(rightNode.Keys, node.Keys[splitIndex+1:]...)

	// Move corresponding children to the new node
	rightNode.Children = append(rightNode.Children, node.Children[splitIndex+1:]...)

	// Update parent pointers for moved children
	for _, child := range rightNode.Children {
		child.Parent = rightNode
	}

	// Update the original node keys and children
	node.Keys = node.Keys[:splitIndex]

	// NOTE: delete underlying Node from ref, for GC purpose.
	for i := splitIndex + 1; i < len(node.Children); i++ {
		// fmt.Printf("disconnect parent %d <-> child %d\n", node.Keys, node.Children[i].Keys)
		node.Children[i] = nil
	}
	node.Children = node.Children[:splitIndex+1]

	return rightNode, promotedKey
}

// InsertIntoParent inserts a key and node into the parent node
func InsertIntoParent(oldLeftNode, newRightNode *Node, key int) *Node {
	// If the node is the root, create a new root
	if oldLeftNode.Parent == nil {
		newRoot := NewNode(false)
		newRoot.Keys = append(newRoot.Keys, key)
		newRoot.Children = append(newRoot.Children, oldLeftNode, newRightNode)
		oldLeftNode.Parent = newRoot
		newRightNode.Parent = newRoot
		return newRoot
	}

	// Otherwise, insert into the oldParent
	oldParent := oldLeftNode.Parent

	// append Key & sort Key
	oldParent.Keys = append(oldParent.Keys, key)
	slices.Sort(oldParent.Keys)

	// append Children & sort Children
	oldParent.Children = append(oldParent.Children, newRightNode)
	slices.SortFunc(oldParent.Children, func(a, b *Node) int {
		// 按照 Key[0] 数值升序
		return a.Keys[0] - b.Keys[0]
	})

	// Set the parent of the new node
	newRightNode.Parent = oldParent

	// If the parent has too many keys, split it
	if len(oldParent.Keys) >= order {
		newParent, promotedKey := SplitInternalNode(oldParent)
		return InsertIntoParent(oldParent, newParent, promotedKey)
	}

	return nil // No new root, insert finished
}
