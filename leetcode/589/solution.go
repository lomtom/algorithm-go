package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:3.8 MB,击败了69.15% 的Go用户
func preorder(root *Node) []int {
	var result []int
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil {
			return
		}
		result = append(result, root.Val)
		for index := range root.Children {
			dfs(root.Children[index])
		}
	}
	dfs(root)
	return result
}
