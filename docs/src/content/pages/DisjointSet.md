---
title: "并查集"
slug: "djs"
---
## 并查集（Disjoint Set Union, DSU）详细介绍

并查集是一种树型的数据结构，用于处理不相交集合（disjoint sets）的合并（union）及查询（find）操作。它特别适合用于动态连通性问题，如网络连接、图的连通分量等。

### 基本操作

- **Find**：查找元素所属集合的代表元素，即可判断两个元素是否属于同一个集合。
- **Union**：合并两个集合。

### 基本概念

- **节点（Node）**：并查集中的每个元素称为一个节点。
- **根（Root）**：一个集合中的代表元素，树的根节点就是该集合的代表元素。
- **路径压缩（Path Compression）**：在执行Find操作时，将沿途访问的节点直接连接到根节点，以加速后续的Find操作。
- **按秩合并（Union by Rank）**：在执行Union操作时，总是将较小的树连接到较大的树上，以保持树的高度尽可能低。

### 例子

假设我们有如下初始并查集，每个元素是一个独立的集合：

```
1   2   3   4   5
```

通过若干次合并操作：

```
Union(1, 2)
Union(3, 4)
Union(2, 4)
```

最终得到：

```
  1
 / \
2   3
     \
      4
```

元素1、2、3、4属于同一个集合。

### 并查集的实现

以下是并查集的基本实现，包含路径压缩和按秩合并优化。

#### 1. 初始化并查集

```go
type DisjointSet struct {
	parent []int
}

func NewDisjointSet(size int) *DisjointSet {
	parent := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
	}
	return &DisjointSet{parent}
}
```

#### 2. 查找操作（Find）

Find 操作用于查找元素所属集合的代表元素。

查询操作Find(x)从x开始，根据节点到父节点的引用向根行进，直到找到根节点。

```go
func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] == x {
		return x
	}
	return ds.Find(ds.parent[x])
}
```

**状态压缩**

在集合很大或者树很不平衡时，上述代码的效率很差，最坏情况下（树退化成一条链时），单次查询的时间复杂度高达 O(𝑛)。那么可以通过状态压缩进行优化


路径压缩优化通过在查找操作过程中，将树的高度尽可能减少，使得后续的查找操作更快。具体方法是，在执行Find操作时，将沿途访问的所有节点直接连接到根节点。

在路径压缩优化中，我们将沿途访问的所有节点直接连接到根节点，以加速后续的查询操作。

```go
func (ds *DisjointSet) Find(x int) int {
    if ds.parent[x] != x {
        ds.parent[x] = ds.Find(ds.parent[x])
    }
    return ds.parent[x]
}
```


例如，查找节点4时，需要首先查找节点3，然后查找节点1，最后才找到根节点1。

执行Find(4)操作前：

```
  1  
 / \
2   3
     \
      4
```
执行Find(4)操作后（路径压缩）：

```
   1
 / | \
2   3  4
```
节点4、3、2都直接连接到根节点1, 这样的话，后续的查询操作都只需要常数时间即可。


#### 3. 合并操作（Union）

合并操作Union(x, y)把元素x所在的集合与元素y所在的集合合并为一个。

合并操作首先找出节点x与节点y对应的两个根节点，如果两个根节点其实是同一个，则说明元素x与元素y已经位于同一个集合中，否则，则使其中一个根节点成为另一个的父节点。

```go
func (ds *DisjointSet) Union(x, y int) {
  rootX := ds.Find(x)
  rootY := ds.Find(y)
  
  if rootX != rootY {
      ds.parent[rootX] = rootY
  }
}

// 或者
func (ds *DisjointSet) Union(x, y int) {
	ds.parent[ds.Find(x)] = ds.Find(y)
}
```

### 复杂度分析

- 时间复杂度几乎接近于常数时间，即 O(𝛼(n))，其中 𝛼(n) 是阿克曼函数的反函数，增长极其缓慢。
- 空间复杂度是 O(𝑛)

### 应用场景

并查集广泛应用于以下场景：
- 动态连通性问题
- 网络连接问题
- 图的连通分量问题
- Kruskal算法中的最小生成树问题

并查集作为一种高效的数据结构，通过简单的操作和强大的优化策略，能在复杂问题中提供高效的解决方案。

### 例题

1. [LeetCode 128. 最长连续序列](/leetcode/longest-consecutive-sequence)
2. [LeetCode 547. 省份数量](/leetcode/number-of-provinces)
2. [LeetCode 990. 等式方程的可满足性](/leetcode/satisfiability-of-equality-equations)
3. [LeetCode 1319. 连通网络的操作次数](/leetcode/number-of-operations-to-make-network-connected)

### 参考资料

- [Wikipedia - Disjoint-set data structure](https://zh.wikipedia.org/wiki/%E5%B9%B6%E6%9F%A5%E9%9B%86)
- [OI Wiki - Disjoint Set Union](https://oi-wiki.org/ds/dsu/#%E6%9F%A5%E8%AF%A2)
