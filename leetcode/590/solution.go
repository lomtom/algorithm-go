package _90

type Node struct {
	Val      int
	Children []*Node
}

// 递归
// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:3.8 MB,击败了51.20% 的Go用户
func postorder(root *Node) []int {
	var result []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for index := range node.Children {
			dfs(node.Children[index])
		}
		result = append(result, node.Val)
	}
	dfs(root)
	return result
}
