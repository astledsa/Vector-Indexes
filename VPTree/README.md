# Vantage Point Tree

One of the first variations on KD Trees, this is perhaps the simplest data structure I could find to index *K*-dimensional vectors. In this file I have implemented the *Insertion* methods for the Index. Following is the construction algorithm of the KD Tree.

## Construction/Insertion

- After initialization, we select a vantage point (root node in a binary tree), which will be used to compare with other nodes.
- During insertion, we get the inner-product between the vantage point and the new node to be inserted (which will be a vector).
- If the inner-product is greater than a pre-determined <i><u>threshold</u></i>, the new node goes to the right, else left
- The inner product could be any of these the methods (Eucilidean Distance, Manhattan Distance, Cosine Similiarity). The <i>threshold</i> value would need to be set accordingly

## Benchmarks

In order to test the Index (and for future comparisons with other indexes), I used the <a href="http://corpus-texmex.irisa.fr/">SIFT1M</a> dataset to benchmark the Index. The details of how I calculated the benchmarks will be below as well. I used <u>10K vectors</b> and <u>10K Queries</u> for the tests (I may add more, but the results seem to be consistent). Since my aim was not to optimise but to learn the data structure itself, I have refrained from using Golang's concurrency in this structure. Although it should be straight forward to do so.

- <b>Index Construction</b> The time it takes to construct (insert) 10,000 128 dimensional vectors. <i>Time Taken : <ins>290.871062s</ins></i>

- <b>Average Search Time</b> The average amount of time it takes for the search algorithm to find the top K queries, where I have set K = 5. I have averaged this value over all the 10,000 queries. <i>Time Taken :  <ins>0.058535s</ins></i>

- <b>Average Precision</b> For each retreived vector, the cosine similiarity is calculated with respect to the original query, and if the similiarity is greater than or equal to 0.5, it's considered a true positive. The true positive value is then divided by the total number of retreived queries (5 in this case) to give the Precision of a single search over the query. These precisions are simply averaged at the end. <i>precision : <ins>0.898310</ins></i>

