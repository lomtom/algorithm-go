package leetcode

func deepestLeavesSum(root *TreeNode) (ans int) {
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		l, temp := len(queue), 0
		for i := 0; i < l; i++ {
			temp += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		ans = temp
		queue = queue[l:]
	}
	return ans
}
