package leetcode

// 执行耗时:2 ms,击败了37.13% 的Go用户
// 内存消耗:2.7 MB,击败了87.01% 的Go用户
func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			pre := curr.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = curr.Right
			curr.Left, curr.Right = nil, curr.Left
		}
		curr = curr.Right
	}
}
