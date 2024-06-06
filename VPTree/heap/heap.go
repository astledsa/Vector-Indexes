package heap

import (
	"errors"
	"fmt"
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
	return &Heap{Root: nil}
}

func (n *Node) Less(o *Node) bool {
	if o == nil {
		return false
	}
	return n.Distance < o.Distance
}

func (n *Node) Greater(o *Node) bool {
	if o == nil {
		return false
	}
	return n.Distance > o.Distance
}

func (n *Node) LessNil(o *Node) bool {
	if o == nil {
		return true
	}
	return n.Distance < o.Distance
}

func (n *Node) GreaterNil(o *Node) bool {
	if o == nil {
		return true
	}
	return n.Distance > o.Distance
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
		Distance:   distance,
		Vector:     vector,
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
	root := tree.Root.Vector

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

	if current.Left != nil {
		current = current.Left
	} else if current.Right != nil {
		current = current.Right
	}

	if current.LeftChild {
		current.Parent.Left = nil
	} else {
		current.Parent.Right = nil
	}

	tree.Root.Distance = current.Distance
	swapped_node := tree.Root
	for swapped_node.Greater(swapped_node.Left) || swapped_node.Greater(swapped_node.Right) {
		if swapped_node.Greater(swapped_node.Left) && swapped_node.Left.LessNil(swapped_node.Right) {
			tempDistance := swapped_node.Left.Distance
			tempVector := swapped_node.Left.Vector
			swapped_node.Left.Distance = swapped_node.Distance
			swapped_node.Left.Vector = swapped_node.Vector
			swapped_node.Distance = tempDistance
			swapped_node.Vector = tempVector
			swapped_node = swapped_node.Left
		} else if swapped_node.Greater(swapped_node.Right) && swapped_node.Right.LessNil(swapped_node.Left) {
			tempDistance := swapped_node.Right.Distance
			tempVector := swapped_node.Right.Vector
			swapped_node.Right.Distance = swapped_node.Distance
			swapped_node.Right.Vector = swapped_node.Vector
			swapped_node.Distance = tempDistance
			swapped_node.Vector = tempVector
			swapped_node = swapped_node.Right
		}
	}

	return root
}

func (tree *Heap) Peek() []float32 {
	return tree.Root.Vector
}

func LeftTraversal(root *Node) {
	if root == nil {
		return
	}
	LeftTraversal(root.Left)
	fmt.Printf("Distance: %f, Vector: %f\n", root.Distance, root.Vector)
}

func RightTraversal(root *Node) {
	if root == nil {
		return
	}
	RightTraversal(root.Right)
	fmt.Printf("Distance: %f, Vector: %f\n", root.Distance, root.Vector)
}

// go run heap.go
