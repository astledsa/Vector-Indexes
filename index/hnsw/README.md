# Heirarchal Navigable Small World Graph

One of the most prominent vector indexes, used among some of the best products is the HNSW graph. It is utilised in Pinecone, Weaviate, Supabase, etc ... and was the best performing index structure for storing meduim to large vectors for a long time. The original paper was written in 2016 of the same name ([here](https://arxiv.org/abs/1603.09320)) and has been dominant since. It is (loosely speaking) a combination of graph and [skiplist](https://en.wikipedia.org/wiki/Skip_list), hence it's a **multi-layered graph**.  

## Construction/Insertion

- The HNSW graph is multi-layered. When a new node is inserted, a random **maximum level** is assigned to it using an exponentially decaying distribution. The node is then added to all layers from level 0 up to its maximum level. The structure maintains an **entry point**, which is the starting node for greedy searches during both insertion and querying.

- During insertion, the algorithm performs a greedy search through the upper layers (from the top down to the target node's maximum level) to locate an approximate closest entry point. Once at or below the node's level, a broader search is performed at each level (usually with `efConstruction > 1`) to find a set of nearest neighbors. These neighbors are then connected to the new node via **bidirectional edges**, subject to a pruning heuristic.

- To search for the K nearest neighbors of a query vector, the algorithm first performs a greedy search from the topmost layer down to layer 1 to find a good starting point. At layer 0, a broader search is performed (typically with `efSearch > K`) to explore the local neighborhood more thoroughly. The top K vectors with the highest similarity are then returned as the result.

## Code Snippet

```Go
func normalizedrandomVector(dim int) la.Vector {
	vec := make([]float32, dim)
	var sumSquares float64
	for i := range vec {
		val := rand.Float64()
		vec[i] = float32(val)
		sumSquares += val * val
	}
	norm := float32(math.Sqrt(sumSquares))
	for i := range vec {
		vec[i] /= norm
	}
	return la.Vector{Data: vec}
}

M := 8
efC := 64
dim := 128
maxLevel := 4
numVectors := 50
ml := 1 / (-math.Log(1 - (1 / float64(M))))

vectors := make([]la.Vector, numVectors)
for i := range vectors {
	vectors[i] = randomVector(dim)
}

index := new(hnsw.HNSW).InitIndex(vectors[0], &M, &efC, &maxLevel)
for i := 1; i < numVectors; i++ {
	index.Insert(vectors[i], maxLevel, ml)
}

query := randomVector(dim)
Results := index.Search(query, 5, 64)
for i, node := range Results {
	fmt.Printf("Result %d: [%.4f ...] (CosSim: %.4f)\n", i+1, node.Data.Data[0], query.CosineSimiliarity(&node.Data))
}
```