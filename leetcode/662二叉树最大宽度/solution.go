package leetcode

import "testing"

// 执行用时：4 ms, 在所有 Go 提交中击败了88.43%的用户
// 内存消耗：4.4 MB, 在所有 Go 提交中击败了60.26%的用户
func widthOfBinaryTree(root *TreeNode) (ans int) {

	type Pair struct {
		*TreeNode
		index int
	}

	queue := []Pair{{
		root,
		1,
	}}

	for len(queue) > 0 {
		l := len(queue)
		left := queue[0].index
		right := left
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, Pair{
					queue[i].Left,
					queue[i].index * 2,
				})
			}
			if queue[i].Right != nil {
				queue = append(queue, Pair{
					queue[i].Right,
					queue[i].index*2 + 1,
				})
			}
			right = queue[i].index
		}
		width := right - left + 1
		if width > ans {
			ans = width
		}
		queue = queue[l:]
	}
	return
}

func TestWidthOfBinaryTree(t *testing.T) {
	widthOfBinaryTree(construct([]int{1, 3, 2, 5}))
}
