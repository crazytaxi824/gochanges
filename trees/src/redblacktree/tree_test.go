package redblacktree

import (
	"fmt"
	"testing"
)

func TestRBtree(*testing.T) {
	tree := NewRBTree()

	// 插入一些值
	tree.Insert(10, "value10")
	tree.Insert(20, "value20")
	tree.Insert(5, "value5")

	// 查找并打印值
	node := tree.Search(10)
	if node != tree.NIL {
		fmt.Println("Found:", node.Value)
	}

	// 中序遍历
	tree.InOrderTraversal(func(node *Node) {
		fmt.Printf("Key: %d, Value: %v\n", node.Key, node.Value)
	})

	// 删除节点
	tree.Delete(10)
}

func TestRBTree2(*testing.T) {
	tree := NewRBTree()

	for i := range 17 {
		tree.Insert(i, nil)
	}

	tree.PrintTree()

	node := tree.Search(10)
	if node != tree.NIL {
		fmt.Println("Found:", node.Key)
	}

	tree.Delete(11)
	tree.PrintTree()
}
