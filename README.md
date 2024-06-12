# Vector-Indexes

A vector index library written in Go, this is purely for my interest in learning the fascinating data structures utilised in the construction of vector indexes. I have decided to build them from scratch in Golang, and test on them personal as well as public datasets. Below are the indexes I've implemented so far, and hopefully will keep adding to them!

## Benchmarks

In order to test the Index (and for future comparisons with other indexes), I used the <a href="http://corpus-texmex.irisa.fr/">SIFT1M</a> dataset to benchmark the Index. I used <u>10K vectors</b> and <u>10K Queries</u> for the tests (I may add more, but the results seem to be consistent). Since my aim was not to optimise but to learn the data structure itself, I have refrained from using Golang's concurrency in this structure. Although it should be straight forward to do so. The details of how I calculated the benchmarks will be below :
- <b>Index Construction</b> The time it takes to construct (insert) 10,000 128 dimensional vectors.
- <b>Average Search Time</b> The average amount of time it takes for the search algorithm to find the top K queries, where I have set K = 5. I have averaged this value over all the 10,000 queries.
- <b>Average Precision</b> For each retreived vector, the cosine similiarity is calculated with respect to the original query, and if the similiarity is greater than or equal to 0.5, it's considered a true positive. The true positive value is then divided by the total number of retreived queries (5 in this case) to give the Precision of a single search over the query. These precisions are simply averaged at the end.

## [1. Vantage Point Tree](https://github.com/Astle-sudo/Vector-Indexes/tree/main/VPTree)

One of the simplest vector index structure, which is just a subset of the KD Trees. These are more or less an extension of the binary tree data structure. Below are the benchmark of my implementation of the structure. For more details, visit the VPTree file. 

<i>Index Construction : <ins>290.871062s</ins></i><br>
<i>Average Search Time:  <ins>0.058535s</ins></i><br>
<i>Precision : <ins>0.898310</ins></i> or <ins>89%</ins>
