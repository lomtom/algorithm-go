package leetcode

// 执行耗时:2 ms,击败了45.38% 的Go用户
// 内存消耗:3.3 MB,击败了88.28% 的Go用户
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	index := l / 2
	return &TreeNode{
		Val:   nums[index],
		Left:  sortedArrayToBST(nums[:index]),
		Right: sortedArrayToBST(nums[index+1:]),
	}
}
