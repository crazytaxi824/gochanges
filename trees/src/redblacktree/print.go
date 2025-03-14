package redblacktree

import (
	"fmt"
	"strings"
)

// PrintTree prints the tree structure to the console
func (t *RBTree) PrintTree() {
	if t.Root == t.NIL {
		fmt.Println("Empty tree")
		return
	}
	t.printTreeRecursive(t.Root, "", true)
}

// printTreeRecursive is a helper function for PrintTree
func (t *RBTree) printTreeRecursive(node *Node, prefix string, isRight bool) {
	if node == t.NIL {
		return
	}

	// Print right subtree first (will appear at the top)
	t.printTreeRecursive(node.Right, prefix+((func() string {
		if isRight {
			return "    "
		}
		return "    "
	})()), false)

	// Print current node
	fmt.Print(prefix)
	if isRight {
		fmt.Print("└── ")
	} else {
		fmt.Print("┌── ")
	}

	// Print with color
	colorName := "[R]"
	if node.Color == BLACK {
		colorName = ""
	}

	// Print node (with color in terminals that support ANSI)
	fmt.Printf("%d%s", node.Key, colorName)

	// Print value if it's a string and not too long
	if str, ok := node.Value.(string); ok && len(str) < 10 {
		fmt.Printf(":%s", str)
	}
	fmt.Println()

	// Print left subtree (will appear at the bottom)
	t.printTreeRecursive(node.Left, prefix+((func() string {
		if isRight {
			return "    "
		}
		return "    "
	})()), true)
}

// PrintTreeSimple prints the tree without ANSI colors (for environments that don't support ANSI)
func (t *RBTree) PrintTreeSimple() {
	if t.Root == t.NIL {
		fmt.Println("Empty tree")
		return
	}
	t.printTreeSimpleRecursive(t.Root, 0)
}

// printTreeSimpleRecursive is a helper function for PrintTreeSimple
func (t *RBTree) printTreeSimpleRecursive(node *Node, level int) {
	if node == t.NIL {
		return
	}

	// Print right subtree
	t.printTreeSimpleRecursive(node.Right, level+1)

	// Print current node
	colorName := "R"
	if node.Color == BLACK {
		colorName = "B"
	}

	fmt.Printf("%s%d[%s]\n", strings.Repeat("    ", level), node.Key, colorName)

	// Print left subtree
	t.printTreeSimpleRecursive(node.Left, level+1)
}

// PrintInOrder prints the tree in order (sorted by key)
func (t *RBTree) PrintInOrder() {
	if t.Root == t.NIL {
		fmt.Println("Empty tree")
		return
	}
	fmt.Print("In-order traversal: ")
	first := true
	t.InOrderTraversal(func(node *Node) {
		if !first {
			fmt.Print(" → ")
		}
		colorName := "R"
		if node.Color == BLACK {
			colorName = "B"
		}
		fmt.Printf("%d[%s]", node.Key, colorName)
		first = false
	})
	fmt.Println()
}
