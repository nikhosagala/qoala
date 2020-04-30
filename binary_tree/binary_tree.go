package binary_tree

type Node struct {
	left  *Node
	right *Node
	value int
}

type BinaryTree struct {
	root *Node
	size int
}

func (bt *BinaryTree) Size() int {
	return bt.size
}

func NewBst() *BinaryTree {
	bst := new(BinaryTree)
	bst.size = 0
	return bst
}

func (root *Node) insert(newNode *Node) {
	if newNode.value > root.value {
		if root.right == nil {
			root.right = newNode
		} else {
			root.right.insert(newNode)
		}
	} else if newNode.value < root.value {
		if root.left == nil {
			root.left = newNode
		} else {
			root.left.insert(newNode)
		}
	}
}

func (bt *BinaryTree) Insert(value int) {
	if bt.root == nil {
		bt.root = &Node{nil, nil, value}
	}
	bt.size++
	bt.root.insert(&Node{nil, nil, value})
}

// i do not understand how to solve this tree problem, sorry
