package leetcode

//执行耗时:4 ms,击败了82.67% 的Go用户
//内存消耗:5.8 MB,击败了10.67% 的Go用户
//func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
//	if depth == 1 {
//		return &TreeNode{
//			Val:  val,
//			Left: root,
//		}
//	}
//	queue := []*TreeNode{root}
//	for depth > 2 {
//		l := len(queue)
//		for index := 0; index < l; index++ {
//			node := queue[index]
//			if node != nil {
//				queue = append(queue, node.Left)
//				queue = append(queue, node.Right)
//			}
//		}
//		queue = queue[l:]
//		depth--
//	}
//	l := len(queue)
//	for index := 0; index < l; index++ {
//		node := queue[index]
//		if node != nil {
//			node.Left = &TreeNode{
//				Val:  val,
//				Left: node.Left,
//			}
//			node.Right = &TreeNode{
//				Val:   val,
//				Right: node.Right,
//			}
//		}
//	}
//	return root
//}

func addOneRow(root *TreeNode, val, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	if depth == 2 {
		root.Left = &TreeNode{val, root.Left, nil}
		root.Right = &TreeNode{val, nil, root.Right}
	} else {
		root.Left = addOneRow(root.Left, val, depth-1)
		root.Right = addOneRow(root.Right, val, depth-1)
	}
	return root
}
