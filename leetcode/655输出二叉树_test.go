package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func printTree(root *TreeNode) [][]string {
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getDepth(node.Left)
		right := getDepth(node.Right)
		if left > right {
			return left + 1
		}
		return right + 1
	}
	// 获取最大深度
	height := getDepth(root)
	ans := make([][]string, height)
	n := 1<<height - 1
	for index := range ans {
		ans[index] = make([]string, n)
	}
	var setValue func(node *TreeNode, r, c int)
	setValue = func(node *TreeNode, r, c int) {
		ans[r][c] = strconv.Itoa(node.Val)
		if node.Left != nil {
			setValue(node.Left, r+1, c-1<<(height-r-2))
		}
		if node.Right != nil {
			setValue(node.Right, r+1, c+1<<(height-r-2))
		}
	}
	setValue(root, 0, (n-1)/2)
	return ans
}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：2.2 MB, 在所有 Go 提交中击败了80.95%
func TestPrintTree(t *testing.T) {
	node := construct([]int{1, 2, 3, 4})
	fmt.Println(printTree(node))
}
