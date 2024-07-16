package leetcode

// 执行耗时:19 ms,击败了40.64% 的Go用户
// 内存消耗:6.4 MB,击败了31.49% 的Go用户
func findCircleNum(isConnected [][]int) int {
	var parent = make([]int, len(isConnected))
	for index := range isConnected {
		parent[index] = index
	}
	var find func(num int) int
	find = func(num int) int {
		if parent[num] != num {
			parent[num] = find(parent[num])
		}
		return parent[num]
	}
	union := func(i, j int) {
		parent[find(j)] = find(i)
	}
	for i := range isConnected {
		for j := i + 1; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 {
				union(i, j)
			}
		}
	}
	var result int
	for index := range parent {
		if parent[index] == index {
			result++
		}
	}
	return result
}

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

func (ds *DisjointSet) Find(x int) int {
	if ds.parent[x] == x {
		return x
	}
	return ds.Find(ds.parent[x])
}

func (ds *DisjointSet) Union(x, y int) {
	ds.parent[ds.Find(x)] = ds.Find(y)
}
