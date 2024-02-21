package _29

type Node struct {
	Val      int
	Children []*Node
}

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:4.2 MB,击败了71.25% 的Go用户
func levelOrder(root *Node) (result [][]int) {
	if root == nil {
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	index := 0
	l := len(queue)
	for len(queue) != index {
		var row = make([]int, 0)
		for ; index < l; index++ {
			row = append(row, queue[index].Val)
			queue = append(queue, queue[index].Children...)
		}
		l = len(queue)
		result = append(result, row)
	}
	return
}
