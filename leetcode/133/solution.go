package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:4.6 MB,击败了30.88% 的Go用户
func cloneGraph(node *Node) *Node {
	var m = make(map[int]*Node)
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, ok := m[node.Val]; ok {
			return n
		}
		n := &Node{Val: node.Val}
		m[node.Val] = n
		for index := range node.Neighbors {
			n.Neighbors = append(n.Neighbors, dfs(node.Neighbors[index]))
		}
		return n
	}
	return dfs(node)
}
