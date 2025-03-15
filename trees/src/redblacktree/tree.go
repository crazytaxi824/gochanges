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
	newNode := &Node{
		Key:    key,
		Value:  value,
		Color:  RED, // 插入节点一开始是红色, 如果有冲突则 fixup 时修改颜色.
		Left:   t.NIL,
		Right:  t.NIL,
		Parent: nil,
	}

	newNodeParent := t.NIL
	x := t.Root

	// Find position for new node from root node.
	for x != t.NIL {
		newNodeParent = x
		if newNode.Key < x.Key {
			x = x.Left
		} else if newNode.Key > x.Key {
			x = x.Right
		} else {
			// Key already exists, update value and return
			x.Value = value
			return
		}
	}

	// Set parent of new node
	// 如果新插入的 node 是 root, 则 root 的 parent 也是 nil.
	newNode.Parent = newNodeParent

	// Insert node
	if newNodeParent == t.NIL {
		t.Root = newNode
	} else if newNode.Key < newNodeParent.Key {
		newNodeParent.Left = newNode
	} else {
		newNodeParent.Right = newNode
	}

	// Fix violations
	t.insertFixup(newNode)
}

// insertFixup fixes violations of red-black tree properties after insertion
func (t *RBTree) insertFixup(newNode *Node) {
	// newNode is not root && newNode Parent is RED
	for newNode.Parent != t.NIL && newNode.Parent.Color == RED {
		if newNode.Parent == newNode.Parent.Parent.Left {
			// newNode Parent is LEFT && newNode Parent is RED
			uncle := newNode.Parent.Parent.Right
			if uncle != t.NIL && uncle.Color == RED {
				// Case 1: Uncle is red
				newNode.Parent.Color = BLACK      // change Parent to BLACK
				uncle.Color = BLACK               // change Uncle to BLACK
				newNode.Parent.Parent.Color = RED // change Grandparent to RED
				newNode = newNode.Parent.Parent   // set Grandparent as a new node
			} else {
				// Case 2: Uncle is NIL or BLACK
				if newNode == newNode.Parent.Right {
					// newNode is right child - inner grandchild
					//      Grandparent
					//        /     \
					//   Parent
					//    /   \
					//       NewNode
					// newNode is Grandparent's "inner grandchild"
					// perform Left-Rotation on Parent Node, then
					// perform Right-Rotation on Grandparent Node in Case 3.
					newNode = newNode.Parent
					t.leftRotate(newNode)
				}
				// Case 3: Uncle is BLACK, newNode is left child - outer grandchild
				newNode.Parent.Color = BLACK         // change Parent to BLACK
				newNode.Parent.Parent.Color = RED    // change Grandparent to BLACK
				t.rightRotate(newNode.Parent.Parent) // perform Right-Rotation on Grandparent Node
			}
		} else {
			// Same cases but mirrored
			uncle := newNode.Parent.Parent.Left
			if uncle != t.NIL && uncle.Color == RED {
				// Case 1: Uncle is red
				newNode.Parent.Color = BLACK
				uncle.Color = BLACK
				newNode.Parent.Parent.Color = RED
				newNode = newNode.Parent.Parent
			} else {
				if newNode == newNode.Parent.Left {
					// Case 2: Uncle is black, newNode is left child - inner grandchild
					// newNode is right child - inner grandchild
					//      Grandparent
					//        /     \
					//             Parent
					//             /   \
					//        NewNode
					// newNode is Grandparent's "inner grandchild"
					// perform Right-Rotation on Parent Node, then
					// perform Left-Rotation on Grandparent Node in Case 3.
					newNode = newNode.Parent
					t.rightRotate(newNode)
				}
				// Case 3: Uncle is black, newNode is right child - outer grandchild
				newNode.Parent.Color = BLACK
				newNode.Parent.Parent.Color = RED
				t.leftRotate(newNode.Parent.Parent)
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

	// 以下两种删除方式都可以
	// t.deleteWithPredecessor(z)
	t.deleteWithSuccessor(z)
}

func (t *RBTree) deleteWithSuccessor(delNode *Node) {
	successor := delNode
	yOriginalColor := successor.Color
	var x *Node

	if delNode.Left == t.NIL {
		// Case 1: delNode has no child OR no left child
		x = delNode.Right
		t.transplant(delNode, delNode.Right) // replace with right child
	} else if delNode.Right == t.NIL {
		// Case 2: delNode has no right child
		x = delNode.Left
		t.transplant(delNode, delNode.Left) // replace with left child
	} else {
		// Case 3: delNode has two children
		// 找到后继节点 successor 代替, 即:右子树中最小的节点, 后继节点没有 left child.
		// 也可以用前驱节点 predecessor 代替, 即:左子树中最大的节点. 但通常使用后继节点代替.
		successor = t.minimumNode(delNode.Right)
		yOriginalColor = successor.Color
		x = successor.Right

		if successor.Parent == delNode {
			// case: successor 是 delNode 的 child
			//  delNode
			//   /   \
			//  A   successor
			//       /    \
			//      NIL    x
			// 'x' could be NIL.
			x.Parent = successor
		} else {
			// case: successor 是 delNode 右子树中最小的节点
			//    delNode
			//     /   \
			//    A     B
			//        /   \
			// successor   ...
			//   /   \
			// NIL    x
			// successor.Right could be NIL.

			// x -> successor
			t.transplant(successor, successor.Right)
			// successor.Right <-> B
			successor.Right = delNode.Right
			successor.Right.Parent = successor
			// 结果:
			//      successor
			//       /    \
			//    ...      B
			//           /   \
			//          x     ...
		}

		// successor -> delNode
		t.transplant(delNode, successor)
		// successor.Left <-> A
		successor.Left = delNode.Left
		successor.Left.Parent = successor
		successor.Color = delNode.Color
		// 结果:
		//      successor
		//       /    \
		//      A      ...
	}

	if yOriginalColor == BLACK {
		t.deleteFixup(x)
	}
}

func (t *RBTree) deleteWithPredecessor(delNode *Node) {
	predecessor := delNode
	yOriginalColor := predecessor.Color
	var x *Node

	if delNode.Left == t.NIL {
		// Case 1: delNode has no child OR no left child
		x = delNode.Right
		t.transplant(delNode, delNode.Right) // replace with right child
	} else if delNode.Right == t.NIL {
		// Case 2: delNode has no right child
		x = delNode.Left
		t.transplant(delNode, delNode.Left) // replace with left child
	} else {
		// Case 3: delNode has two children
		// 找到前驱节点 predecessor 代替, 即:左子树中最大的节点, 后继节点没有 right child.
		predecessor = t.maximumNode(delNode.Left)
		yOriginalColor = predecessor.Color
		x = predecessor.Right

		if predecessor.Parent == delNode {
			// case: predecessor 是 delNode 的 child
			x.Parent = predecessor
		} else {
			// case: predecessor 是 delNode 的左子树中最大的节点.
			t.transplant(predecessor, predecessor.Left)
			predecessor.Left = delNode.Left
			predecessor.Left.Parent = predecessor
		}
		t.transplant(delNode, predecessor)
		predecessor.Right = delNode.Right
		predecessor.Right.Parent = predecessor
		predecessor.Color = delNode.Color
	}

	if yOriginalColor == BLACK {
		t.deleteFixup(x)
	}
}

// deleteFixup fixes violations of red-black tree properties after deletion
func (t *RBTree) deleteFixup(x *Node) {
	for x != t.Root && x.Color == BLACK {
		if x == x.Parent.Left {
			// x is left side
			sibling := x.Parent.Right
			if sibling.Color == RED {
				// Case 1: x's sibling w is RED
				sibling.Color = BLACK
				x.Parent.Color = RED
				t.leftRotate(x.Parent) // x is left perform Left-Rotation; x is right perform Right-Rotation
				sibling = x.Parent.Right
			}
			if sibling.Left.Color == BLACK && sibling.Right.Color == BLACK {
				// Case 2: Both of w's children are BLACK
				sibling.Color = RED
				x = x.Parent // x -> x.Parent
			} else {
				if sibling.Right.Color == BLACK {
					// Case 3: w's right child is BLACK, outer child is BLACK, inner child is RED
					sibling.Left.Color = BLACK
					sibling.Color = RED
					t.rightRotate(sibling)
					sibling = x.Parent.Right
				}
				// Case 4: w's right child is RED, outer child is RED, inner child is BLACK
				sibling.Color = x.Parent.Color
				x.Parent.Color = BLACK
				sibling.Right.Color = BLACK
				t.leftRotate(x.Parent) // x is left perform Left-Rotation; x is right perform Right-Rotation
				x = t.Root
			}
		} else {
			// x is right side, Same cases but mirrored
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

// replaces one subtree 'u' with another 'v'
func (t *RBTree) transplant(u, v *Node) {
	if u.Parent == t.NIL {
		// u is Root
		t.Root = v
	} else if u == u.Parent.Left {
		// u is LEFT
		u.Parent.Left = v
	} else {
		// u is RIGHT
		u.Parent.Right = v
	}
	v.Parent = u.Parent
}

// minimumNode finds the node with minimumNode key in the subtree rooted at x
// search for successor
func (t *RBTree) minimumNode(x *Node) *Node {
	// left child < parent < right child
	for x.Left != t.NIL {
		x = x.Left
	}
	return x
}

// search for predecessor
func (t *RBTree) maximumNode(x *Node) *Node {
	// left child < parent < right child
	for x.Right != t.NIL {
		x = x.Right
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
