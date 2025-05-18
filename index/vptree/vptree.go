package vptree

import (
	"math"

	"VIndex/common/heap"
	"VIndex/common/la"
)

type Node struct {
	Left      *Node
	Right     *Node
	Vector    *la.Vector
	Threshold float32
}

type Tree struct {
	totalNodes int
	Root       *Node
}

func Init_VPTree() *Tree {
	new_tree := Tree{Root: nil, totalNodes: 0}
	return &new_tree
}

func (tree *Tree) Insert(kd_point []float32) {
	new_node := Node{
		Left:      nil,
		Right:     nil,
		Vector:    &la.Vector{Data: kd_point},
		Threshold: 0,
	}

	if tree.Root == nil {
		tree.Root = &new_node
		return
	}

	current := tree.Root
	for current.Left != nil || current.Right != nil {
		if new_node.Vector.CosineSimiliarity(current.Vector) > current.Threshold {
			current = current.Right
		} else {
			current = current.Left
		}
	}
	if current.Vector.CosineSimiliarity(new_node.Vector) > current.Threshold {
		tree.totalNodes += 1
		current.Right = &new_node
	} else {
		tree.totalNodes += 1
		current.Left = &new_node
	}
}

func (tree *Tree) Search(query []float32, K int) *heap.Heap {
	queryVector := la.Vector{Data: query}
	priorityQueue := heap.BinaryHeap()

	current := tree.Root
	priorityQueue.Insert(queryVector.CosineSimiliarity(current.Vector), current.Vector.Data)
	for current != nil {
		distance := queryVector.CosineSimiliarity(current.Vector)
		if math.Abs(float64(distance)) > 0.5 {
			priorityQueue.Insert(distance, current.Vector.Data)
		}
		if distance > current.Threshold {
			current = current.Right
		} else {
			current = current.Left
		}
	}

	return priorityQueue
}
