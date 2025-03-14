package bplustree

import (
	"testing"
)

func TestInsertFind(t *testing.T) {
	tree := NewBPlusTree()

	for i := range 13 {
		err := tree.Insert(i)
		if err != nil {
			t.Error(err)
			return
		}
	}

	PrintTree(tree.Root, 0)

	// test internal function
	t.Log(tree.findLeafNode(3).Keys)
	t.Log(tree.findLeafNode(11).Keys)

	// test Search
	s, _ := tree.Search(3)
	t.Logf("%+v", s)

	s, _ = tree.Search(11)
	t.Logf("%+v", s)
}
