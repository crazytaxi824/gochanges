package redblacktree

// Color represents the color of a node in the red-black tree
type Color bool

const (
	RED   Color = true
	BLACK Color = false
)

// Node represents a node in the red-black tree
type Node struct {
	Key    int
	Value  any
	Color  Color
	Left   *Node
	Right  *Node
	Parent *Node
}

// RBTree represents a red-black tree
type RBTree struct {
	Root *Node
	NIL  *Node // Sentinel node, 哨兵节点
}

// NewRBTree creates a new red-black tree
func NewRBTree() *RBTree {
	nil_node := &Node{Color: BLACK}
	return &RBTree{
		NIL:  nil_node,
		Root: nil_node,
	}
}

// Search finds a node with the given key
func (t *RBTree) Search(key int) *Node {
	return t.search(t.Root, key)
}

func (t *RBTree) search(x *Node, key int) *Node {
	if x == t.NIL || x.Key == key {
		return x
	}
	if key < x.Key {
		return t.search(x.Left, key)
	}
	return t.search(x.Right, key)
}

// Insert adds a new node with the given key and value
func (t *RBTree) Insert(key int, value any) {
	// Create new node
	z := &Node{
		Key:    key,
		Value:  value,
		Color:  RED,
		Left:   t.NIL,
		Right:  t.NIL,
		Parent: nil,
	}

	y := t.NIL
	x := t.Root

	// Find position for new node
	for x != t.NIL {
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else if z.Key > x.Key {
			x = x.Right
		} else {
			// Key already exists, update value and return
			x.Value = value
			return
		}
	}

	// Set parent of z
	z.Parent = y

	// Insert node
	if y == t.NIL {
		t.Root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}

	// Fix violations
	t.insertFixup(z)
}

// insertFixup fixes violations of red-black tree properties after insertion
func (t *RBTree) insertFixup(z *Node) {
	for z.Parent != t.NIL && z.Parent.Color == RED {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y != t.NIL && y.Color == RED {
				// Case 1: Uncle is red
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					// Case 2: Uncle is black, z is right child
					z = z.Parent
					t.leftRotate(z)
				}
				// Case 3: Uncle is black, z is left child
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.rightRotate(z.Parent.Parent)
			}
		} else {
			// Same cases but mirrored
			y := z.Parent.Parent.Left
			if y != t.NIL && y.Color == RED {
				// Case 1: Uncle is red
				z.Parent.Color = BLACK
				y.Color = BLACK
				z.Parent.Parent.Color = RED
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					// Case 2: Uncle is black, z is left child
					z = z.Parent
					t.rightRotate(z)
				}
				// Case 3: Uncle is black, z is right child
				z.Parent.Color = BLACK
				z.Parent.Parent.Color = RED
				t.leftRotate(z.Parent.Parent)
			}
		}
	}
	t.Root.Color = BLACK
}

// Delete removes a node with the given key
func (t *RBTree) Delete(key int) {
	z := t.Search(key)
	if z == t.NIL {
		return
	}

	t.delete(z)
}

func (t *RBTree) delete(z *Node) {
	y := z
	yOriginalColor := y.Color
	var x *Node

	if z.Left == t.NIL {
		// Case 1: z has no left child
		x = z.Right
		t.transplant(z, z.Right)
	} else if z.Right == t.NIL {
		// Case 2: z has no right child
		x = z.Left
		t.transplant(z, z.Left)
	} else {
		// Case 3: z has two children
		y = t.minimum(z.Right)
		yOriginalColor = y.Color
		x = y.Right

		if y.Parent == z {
			x.Parent = y
		} else {
			t.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}

		t.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
		y.Color = z.Color
	}

	if yOriginalColor == BLACK {
		t.deleteFixup(x)
	}
}

// deleteFixup fixes violations of red-black tree properties after deletion
func (t *RBTree) deleteFixup(x *Node) {
	for x != t.Root && x.Color == BLACK {
		if x == x.Parent.Left { // x is left side
			w := x.Parent.Right
			if w.Color == RED {
				// Case 1: x's sibling w is RED
				w.Color = BLACK
				x.Parent.Color = RED
				t.leftRotate(x.Parent)
				w = x.Parent.Right
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				// Case 2: Both of w's children are BLACK
				w.Color = RED
				x = x.Parent
			} else {
				if w.Right.Color == BLACK {
					// Case 3: w's right child is BLACK
					w.Left.Color = BLACK
					w.Color = RED
					t.rightRotate(w)
					w = x.Parent.Right
				}
				// Case 4: w's right child is RED
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Right.Color = BLACK
				t.leftRotate(x.Parent)
				x = t.Root
			}
		} else { // x is right side
			// Same cases but mirrored
			w := x.Parent.Left
			if w.Color == RED {
				// Case 1: x's sibling w is RED
				w.Color = BLACK
				x.Parent.Color = RED
				t.rightRotate(x.Parent)
				w = x.Parent.Left
			}
			if w.Right.Color == BLACK && w.Left.Color == BLACK {
				// Case 2: Both of w's children are BLACK
				w.Color = RED
				x = x.Parent
			} else {
				if w.Left.Color == BLACK {
					// Case 3: w's left child is BLACK
					w.Right.Color = BLACK
					w.Color = RED
					t.leftRotate(w)
					w = x.Parent.Left
				}
				// Case 4: w's left child is RED
				w.Color = x.Parent.Color
				x.Parent.Color = BLACK
				w.Left.Color = BLACK
				t.rightRotate(x.Parent)
				x = t.Root
			}
		}
	}
	x.Color = BLACK
}

// leftRotate performs a left rotation on the given node
func (t *RBTree) leftRotate(x *Node) {
	y := x.Right
	x.Right = y.Left

	if y.Left != t.NIL {
		y.Left.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == t.NIL {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

// rightRotate performs a right rotation on the given node
func (t *RBTree) rightRotate(y *Node) {
	x := y.Left
	y.Left = x.Right

	if x.Right != t.NIL {
		x.Right.Parent = y
	}

	x.Parent = y.Parent

	if y.Parent == t.NIL {
		t.Root = x
	} else if y == y.Parent.Left {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}

	x.Right = y
	y.Parent = x
}

// transplant replaces one subtree with another
func (t *RBTree) transplant(u, v *Node) {
	if u.Parent == t.NIL {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}

// minimum finds the node with minimum key in the subtree rooted at x
func (t *RBTree) minimum(x *Node) *Node {
	for x.Left != t.NIL {
		x = x.Left
	}
	return x
}

// InOrderTraversal traverses the tree in-order and executes the given function for each node
func (t *RBTree) InOrderTraversal(fn func(*Node)) {
	t.inOrderTraversal(t.Root, fn)
}

func (t *RBTree) inOrderTraversal(x *Node, fn func(*Node)) {
	if x != t.NIL {
		t.inOrderTraversal(x.Left, fn)
		fn(x)
		t.inOrderTraversal(x.Right, fn)
	}
}
