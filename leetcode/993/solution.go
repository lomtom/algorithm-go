package leetcode

// 执行耗时:2 ms,击败了12.28% 的Go用户
// 内存消耗:2.3 MB,击败了36.85% 的Go用户
func isCousins(root *TreeNode, x int, y int) bool {
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		var queue1 []*TreeNode
		var flagX, flagY = false, false
		for index := range queue {
			if (queue[index].Left != nil && queue[index].Right != nil) && ((queue[index].Left.Val == x && queue[index].Right.Val == y) ||
				(queue[index].Left.Val == y && queue[index].Right.Val == x)) {
				return false
			}
			if queue[index].Left != nil {
				queue1 = append(queue1, queue[index].Left)
			}
			if queue[index].Right != nil {
				queue1 = append(queue1, queue[index].Right)
			}
			if x == queue[index].Val {
				flagX = true
			}
			if y == queue[index].Val {
				flagY = true
			}
		}
		if flagX && flagY {
			return true
		}
		queue = queue1
	}
	return false
}
