package leetcode

// 执行耗时:346 ms,击败了34.38% 的Go用户
// 内存消耗:41.6 MB,击败了81.25% 的Go用户
func closestNodes(root *TreeNode, queries []int) [][]int {
	var nums []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		nums = append(nums, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	midSearch := func(x int) int {
		left := 0
		right := len(nums)
		for left < right {
			mid := left + (right-left)/2
			if nums[mid] >= x {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return left
	}
	result := make([][]int, len(queries))
	for index := range queries {
		min, max := -1, -1
		searchIndex := midSearch(queries[index])
		if searchIndex < len(nums) {
			max = nums[searchIndex]
			if nums[searchIndex] == queries[index] {
				min = nums[searchIndex]
				result[index] = []int{min, max}
				continue
			}
		}
		if searchIndex != 0 {
			min = nums[searchIndex-1]
		}
		result[index] = []int{min, max}
	}
	return result
}
