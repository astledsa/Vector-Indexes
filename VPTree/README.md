# Vantage Point Tree

One of the first variations on KD Trees, this is perhaps the simplest data structure I could find to index *K*-dimensional vectors. In this file I have implemented the *Insertion* methods for the Index. Following is the construction algorithm of the KD Tree.

## Construction/Insertion

- After initialization, we select a vantage point (root node in a bianry tree), which will be used to compare with other nodes.
- During insertion, we get the inner-product between the vantage point and the new node to be inserted (which will be a vector).
- If the inner-product is greater than a pre-determined <i><u>threshold</u></i>, the new node goes to the right, else left

## Benchmarks

In order to test the Index (and for future comparisons with other indexes), I used the <a href="http://corpus-texmex.irisa.fr/">SIFT1M</a> dataset to benchmark the Index. The details of how I calculated the benchmarks will be below as well. I used <u>10K vectors</b> and <u>10K Queries</u> for the tests (I may add more, but the results seem to be consistent). Since my aim was not to optimise but to learn the data structure itself, I have refrained from using Golang's concurrency in this structure. Although it should be straight forward to do so.

<i>Index Construction</i><br>
<i>Average Search Time</i><br>
<i>Average Precision</i><br>

