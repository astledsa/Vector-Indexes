package vptree

import (
	"math"
	"math/rand"

	"VPTree/heap"
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
	Root *Node
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

func Vantage_Point(kd_point []float32) *Tree {
	new_vector := Vector{
		Data: kd_point,
	}
	new_node := Node{
		Left:      nil,
		Right:     nil,
		Vector:    &new_vector,
		Threshold: 0,
	}
	new_tree := Tree{
		Root: &new_node,
	}

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
		current.Right = &new_node
	} else {
		current.Left = &new_node
	}
}

func (tree *Tree) Search(query []float32, K int) *heap.Heap {
	queryVector := Vector{Data: query}
	priorityQueue := heap.BinaryHeap()

	current := tree.Root
	priorityQueue.Insert(queryVector.CosineSimiliarity(current.Vector), current.Vector.Data)
	for (current != nil) && (priorityQueue.Root.RightNodes+priorityQueue.Root.LeftNodes != K) {
		distance := queryVector.CosineSimiliarity(current.Vector)
		priorityQueue.Insert(distance, current.Vector.Data)
		if distance > current.Threshold {
			current = current.Right
		} else {
			current = current.Left
		}
	}

	return priorityQueue
}

func GenerateRandomFloat32Array(K int) []float32 {
	array := make([]float32, K)
	for i := range array {
		array[i] = rand.Float32()
	}
	return array
}

// func main() {
// 	K := 10
// 	tree := Vantage_Point(GenerateRandomFloat32Array(K))
// 	for i := 0; i < 100; i++ {
// 		tree.Insert(GenerateRandomFloat32Array(K))
// 	}

// 	query := GenerateRandomFloat32Array(K)
// 	Q := tree.Search(query, 5)

// 	heap.LeftTraversal(Q.Root)

// }

// go run main.go
