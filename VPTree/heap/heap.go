package heap

import (
	"errors"
)

type Node struct {
	Left       *Node
	Right      *Node
	Parent     *Node
	RightNodes int
	LeftNodes  int
	LeftChild  bool
	RightChild bool
	Distance   float32
	Vector     []float32
}

type Heap struct {
	Root *Node
}

func BinaryHeap() *Heap {
	return &Heap{
		Root: nil,
	}
}

func (n *Node) Less(o *Node) bool {
	if o == nil {
		return false
	}
	return n.Distance < o.Distance
}

func (n *Node) LessNil(o *Node) bool {
	if o == nil {
		return true
	}
	return n.Distance < o.Distance
}

func (n *Node) Greater(o *Node) bool {
	if o == nil {
		return false
	}
	return n.Distance > o.Distance
}

func SwapNodes(node1 *Node, node2 *Node) {
	tempDistance := node1.Distance
	tempVector := node1.Vector
	node1.Distance = node2.Distance
	node1.Vector = node2.Vector
	node2.Distance = tempDistance
	node2.Vector = tempVector
}

func Recount(leaf *Node) {
	if leaf.Parent == nil {
		return
	}
	if leaf.LeftChild {
		leaf.Parent.LeftNodes -= 1
	} else if leaf.RightChild {
		leaf.Parent.RightNodes -= 1
	}

	Recount(leaf.Parent)
}

func (tree *Heap) Insert(distance float32, vector []float32) error {
	node_to_insert := Node{
		Left:       nil,
		Right:      nil,
		Parent:     nil,
		RightNodes: 0,
		LeftNodes:  0,
		LeftChild:  false,
		RightChild: false,

		Distance: distance,
		Vector:   vector,
	}

	if tree.Root == nil {
		tree.Root = &node_to_insert
		return nil
	}

	current := tree.Root
	for current.Left != nil && current.Right != nil {
		if current.LeftNodes == current.RightNodes {
			current.LeftNodes += 1
			current = current.Left
		} else {
			current.RightNodes += 1
			current = current.Right
		}
	}

	node_to_insert.Parent = current
	if current.Left == nil {
		current.LeftNodes += 1
		node_to_insert.LeftChild = true
		current.Left = &node_to_insert
	} else if current.Right == nil {
		current.RightNodes += 1
		node_to_insert.RightChild = true
		current.Right = &node_to_insert
	} else {
		return errors.New("could not insert item")
	}

	new_node := &node_to_insert
	for new_node.Less(new_node.Parent) {
		tempDistance := new_node.Parent.Distance
		tempVector := new_node.Parent.Vector
		new_node.Parent.Distance = new_node.Distance
		new_node.Parent.Vector = new_node.Vector
		new_node.Distance = tempDistance
		new_node.Vector = tempVector
		new_node = new_node.Parent
	}

	return nil
}

func (tree *Heap) Pop() []float32 {
	root_vector := tree.Root.Vector

	last_node := tree.Root
	for {
		if last_node.Right == nil && last_node.Left == nil {
			break
		} else {
			if last_node.LeftNodes == last_node.RightNodes {
				last_node = last_node.Right
			} else {
				last_node = last_node.Left
			}
		}
	}

	SwapNodes(last_node, tree.Root)
	Recount(last_node)
	if last_node.Parent != nil {
		if last_node.LeftChild {
			last_node.Parent.Left = nil
		} else {
			last_node.Parent.Right = nil
		}
	}

	node_to_swap := tree.Root
	for (node_to_swap.Greater(node_to_swap.Left) || node_to_swap.Greater(node_to_swap.Right)) &&
		(node_to_swap.Right != nil && node_to_swap.Left != nil) {
		if node_to_swap.Left.LessNil(node_to_swap.Right) {
			SwapNodes(node_to_swap, node_to_swap.Left)
			node_to_swap = node_to_swap.Left
		} else if node_to_swap.Right != nil {
			SwapNodes(node_to_swap, node_to_swap.Right)
			node_to_swap = node_to_swap.Right
		}
	}

	return root_vector
}

func (tree *Heap) Peek() []float32 {
	return tree.Root.Vector
}

// go run heap.go
