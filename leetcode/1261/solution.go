package leetcode

// 执行耗时:26 ms,击败了59.26% 的Go用户
// 内存消耗:8.1 MB,击败了14.81% 的Go用户

type FindElements struct {
	m map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	f := FindElements{m: make(map[int]bool)}
	root.Val = 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			f.m[node.Val] = true
			if node.Left != nil {
				node.Left.Val = 2*node.Val + 1
			}
			if node.Right != nil {
				node.Right.Val = 2*node.Val + 2
			}
			dfs(node.Left)
			dfs(node.Right)
		}
	}
	dfs(root)
	return f
}

func (e *FindElements) Find(target int) bool {
	return e.m[target]
}
