package leetcode

import (
	"math"
	"sort"
)

func verticalTraversal(root *TreeNode) (ans [][]int) {
	type Pair struct {
		val, row, col int
	}
	var data []Pair
	var dfs func(node *TreeNode, row int, col int)
	dfs = func(node *TreeNode, row int, col int) {
		if node != nil {
			data = append(data, Pair{
				node.Val, row, col,
			})
			dfs(node.Left, row+1, col-1)
			dfs(node.Right, row+1, col+1)
		}
	}
	dfs(root, 0, 0)
	sort.Slice(data, func(i, j int) bool {
		return data[i].col < data[j].col || data[i].col == data[j].col && (data[i].row < data[j].row || data[i].row == data[j].row && data[i].val < data[j].val)
	})

	lastCol := math.MinInt32
	for _, node := range data {
		if node.col != lastCol {
			lastCol = node.col
			ans = append(ans, nil)
		}
		ans[len(ans)-1] = append(ans[len(ans)-1], node.val)
	}
	return
}
