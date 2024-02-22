package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.1 MB,击败了55.24% 的Go用户
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	var queue = make([]*TreeNode, 0)
	var index = 0
	queue = append(queue, root)
	for len(queue) != 0 {
		var queue1 = make([]*TreeNode, 0)
		ans = append(ans, queue[0].Val)
		for i := range queue {
			ans[index] = queue[i].Val
			if queue[i].Left != nil {
				queue1 = append(queue1, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue1 = append(queue1, queue[i].Right)
			}
		}
		index++
		queue = queue1
	}
	return
}
