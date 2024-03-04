package leetcode

// 执行耗时:317 ms,击败了12.67% 的Go用户
// 内存消耗:37.8 MB,击败了30.98% 的Go用户
func reachableNodes(n int, edges [][]int, restricted []int) int {
	edgesMap := make(map[int][]int)
	visitedMap := make(map[int]bool)
	for index := range edges {
		edgesMap[edges[index][0]] = append(edgesMap[edges[index][0]], edges[index][1])
		edgesMap[edges[index][1]] = append(edgesMap[edges[index][1]], edges[index][0])
	}
	for index := range restricted {
		visitedMap[restricted[index]] = true
	}
	visitedMap[0] = true
	var count int
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) != 0 {
		num := queue[0]
		queue = queue[1:]
		for index := range edgesMap[num] {
			if !visitedMap[edgesMap[num][index]] {
				visitedMap[edgesMap[num][index]] = true
				count++
				queue = append(queue, edgesMap[num][index])
			}
		}
	}
	return count + 1
}
