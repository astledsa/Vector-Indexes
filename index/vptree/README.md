# Vantage Point Tree

One of the first variations on KD Trees, this is perhaps the simplest data structure I could find to index *K*-dimensional vectors. In this file I have implemented the *Insertion* methods for the Index. Following is the construction algorithm of the KD Tree.

## Construction/Insertion

- After initialization, we select a vantage point (root node in a binary tree), which will be used to compare with other nodes.
- During insertion, we get the inner-product between the vantage point and the new node to be inserted (which will be a vector).
- If the inner-product is greater than a pre-determined <i><u>threshold</u></i>, the new node goes to the right, else left
- The inner product could be any of these methods (Eucilidean Distance, Manhattan Distance, Cosine Similiarity). The <i>threshold</i> value would need to be set accordingly. In this implementation, I have used the cosine similiarity, and used a threshold value of 0.

## Code Snippet

```Go
func randomFloat32Array(K int) []float32 {
	result := make([]float32, K)
	for i := range result {
		result[i] = rand.Float32()
	}
	return result
}

func main() {
	K := 128
	number_of_vectors_to_index := 100
	vantage_point_tree := vptree.Init_VPTree()

	for i := 0; i < number_of_vectors_to_index; i++ {
		vantage_point_tree.Insert(randomFloat32Array(K))
	}

	query := randomFloat32Array(K)
	priorityQueue := vantage_point_tree.Search(query, 5)
	fmt.Println(priorityQueue.Pop())
}
```



 
