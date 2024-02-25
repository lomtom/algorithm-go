package leetcode

import "sort"

// 执行耗时:206 ms,击败了18.68% 的Go用户
// 内存消耗:26.71 MB,击败了22.50% 的Go用户
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	result := make([]int, 0)
	for len(queue) != 0 {
		queue1 := make([]*TreeNode, 0)
		var sum = 0
		for index := range queue {
			if queue[index] == nil {
				continue
			}
			sum += queue[index].Val
			queue1 = append(queue1, queue[index].Left)
			queue1 = append(queue1, queue[index].Right)
		}
		queue = queue1
		result = append(result, sum)
	}
	if len(result) <= k {
		return -1
	}
	sort.Ints(result)
	return int64(result[len(result)-k])
}
