package vptree

import (
	"math"

	"test/indexes/VPTree/heap"
)

type Vector struct {
	Data []float32
}

type Node struct {
	Left      *Node
	Right     *Node
	Vector    *Vector
	Threshold float32
}

type Tree struct {
	totalNodes int
	Root       *Node
}

func (vector *Vector) L1() float32 {
	var sum float64

	for _, i := range vector.Data {
		sum += math.Abs(float64(i))
	}

	return float32(sum)
}

func (vector *Vector) L2() float32 {
	var sum float64

	for _, i := range vector.Data {
		sum += math.Pow(float64(i), 2)
	}

	return float32(math.Sqrt(sum))
}

func (v *Vector) Dot(o *Vector) float32 {
	var sum float32
	for i := range v.Data {
		sum += v.Data[i] * o.Data[i]
	}
	return sum
}

func (v *Vector) CosineSimiliarity(o *Vector) float32 {
	return v.Dot(o) / (v.L2() * o.L2())
}

func Init_VPTree() *Tree {
	new_tree := Tree{Root: nil, totalNodes: 0}
	return &new_tree
}

func (tree *Tree) Insert(kd_point []float32) {
	new_node := Node{
		Left:      nil,
		Right:     nil,
		Vector:    &Vector{Data: kd_point},
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
	queryVector := Vector{Data: query}
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
