package hnsw

import (
	"VIndex/common/la"
	"math"
	"math/rand/v2"
	"slices"
	"sort"
)

type Node struct {
	Max_level  int
	Data       la.Vector
	Neighbours map[int][]*Node
}

type HNSW struct {
	M              int
	Ep             *Node
	Total_nodes    int
	Efconstruction int
}

type Nodes []*Node

func IntOrDefault(p *int, def int) int {
	if p != nil {
		return *p
	}
	return def
}

func Append(nodes *Nodes, newNode Node) {
	if nodes.Includes(&newNode) {
		return
	}

	*nodes = append(*nodes, &newNode)
}

func AddNeighbour(node1, node2 *Node, layer int) {
	neighbours := node1.Neighbours[layer]
	for _, node := range neighbours {
		if node.Data.Equal(&node2.Data) {
			return
		}
	}
	node1.Neighbours[layer] = append(node1.Neighbours[layer], node2)
}

func (nodes Nodes) Includes(q *Node) bool {
	for _, node := range nodes {
		if node.Data.Equal(&q.Data) {
			return true
		}
	}
	return false
}

func (nodes Nodes) RemoveNode(target *Node) Nodes {
	for i, n := range nodes {
		if n == target {
			return slices.Delete(nodes, i, i+1)
		}
	}
	return nodes
}

func (nodes Nodes) Sort(q la.Vector) {
	sort.Slice(nodes, func(i, j int) bool {
		return q.CosineSimiliarity(&nodes[i].Data) > q.CosineSimiliarity(&nodes[j].Data)
	})
}

func (hnsw *HNSW) InitIndex(point la.Vector, M, efC, MAX_LEVEL *int) *HNSW {

	m := IntOrDefault(M, 5)
	efc := IntOrDefault(efC, 100)
	max_level := IntOrDefault(MAX_LEVEL, 5)
	ml := 1 / (-math.Log(1 - (1 / float64(m))))
	node_level := min(int(-math.Log(rand.Float64())*ml), max_level)

	newNode := &Node{
		Data:       point,
		Max_level:  node_level,
		Neighbours: make(map[int][]*Node),
	}
	for i := range node_level {
		newNode.Neighbours[i] = []*Node{}
	}

	return &HNSW{
		M:              m,
		Ep:             newNode,
		Total_nodes:    1,
		Efconstruction: efc,
	}
}

func Nearest(vectors Nodes, element la.Vector) *Node {
	if len(vectors) == 0 {
		return nil
	}

	nearest := vectors[0]
	maxSim := vectors[0].Data.CosineSimiliarity(&element)

	for _, node := range vectors[1:] {
		sim := node.Data.CosineSimiliarity(&element)
		if sim > maxSim {
			maxSim = sim
			nearest = node
		}
	}

	return nearest
}

func Furthest(vectors Nodes, element la.Vector) *Node {
	if len(vectors) == 0 {
		return nil
	}

	nearest := vectors[0]
	minSim := vectors[0].Data.CosineSimiliarity(&element)

	for _, node := range vectors[1:] {
		sim := node.Data.CosineSimiliarity(&element)
		if sim < minSim {
			minSim = sim
			nearest = node
		}
	}

	return nearest
}

func SearchLayer(q la.Vector, ep *Node, ef, layer int) Nodes {

	visited := Nodes{ep}
	candidates := Nodes{ep}
	neighbours := Nodes{ep}

	for len(candidates) > 0 {
		c := Nearest(candidates, q)
		f := Furthest(neighbours, q)

		for _, e := range c.Neighbours[layer] {
			if !visited.Includes(e) {
				Append(&visited, *e)
				f := Furthest(neighbours, q)
				if q.CosineSimiliarity(&e.Data) > q.CosineSimiliarity(&f.Data) || len(neighbours) < ef {
					Append(&candidates, *e)
					Append(&neighbours, *e)

					if len(neighbours) > ef {
						f := Furthest(neighbours, q)
						neighbours = neighbours.RemoveNode(f)
					}
				}
			}
		}

		if q.CosineSimiliarity(&c.Data) < q.CosineSimiliarity(&f.Data) {
			break
		}

		candidates = candidates.RemoveNode(c)
	}

	return neighbours
}

func SelectNeighbours(q la.Vector, C Nodes, M int) Nodes {
	C.Sort(q)
	if len(C) > M {
		return C[:M]
	}
	return C
}

func BidirectionalConnection(nodes Nodes, newNode *Node, M, layer int) {
	for _, node := range nodes {

		AddNeighbour(node, newNode, layer)
		AddNeighbour(newNode, node, layer)
		pruned := SelectNeighbours(node.Data, node.Neighbours[layer], M)
		node.Neighbours[layer] = pruned

	}

	newNode.Neighbours[layer] = SelectNeighbours(newNode.Data, newNode.Neighbours[layer], M)
}

func (hnsw *HNSW) Insert(q la.Vector, max_level int, ml float64) {

	ep := hnsw.Ep
	L := hnsw.Ep.Max_level
	l := min(int(-math.Log(rand.Float64())*ml), max_level)

	newNode := &Node{
		Max_level:  l,
		Data:       q,
		Neighbours: make(map[int][]*Node),
	}

	for i := range l {
		newNode.Neighbours[i] = []*Node{}
	}

	for l_c := L; l_c >= l; l_c-- {
		W := SearchLayer(q, ep, 1, l_c)
		ep = Nearest(W, q)
	}

	for l_c := min(L, l); l_c >= 0; l_c-- {
		W := SearchLayer(q, ep, hnsw.Efconstruction, l_c)
		neighbours := SelectNeighbours(q, W, hnsw.M)
		BidirectionalConnection(neighbours, newNode, hnsw.M, l_c)
	}

	if l > L {
		hnsw.Ep = newNode
	}

	hnsw.Total_nodes += 1
}

func (hnsw *HNSW) Search(q la.Vector, K, efSearch int) Nodes {

	ep := hnsw.Ep
	L := hnsw.Ep.Max_level

	for l_c := L; l_c >= 1; l_c-- {
		W := SearchLayer(q, ep, 1, l_c)
		ep = Nearest(W, q)
	}

	nearest := SearchLayer(q, ep, efSearch, 0)
	nearest.Sort(q)

	if len(nearest) > K {
		return nearest[:K]
	}

	return nearest
}
