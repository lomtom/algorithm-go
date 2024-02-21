package leetcode

// 输入：inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 输出：[3,9,20,null,null,15,7]

// 执行耗时:6 ms,击败了26.59% 的Go用户
// 内存消耗:3.6 MB,击败了93.35% 的Go用户
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	var index int
	for index = range inorder {
		if inorder[index] == postorder[len(postorder)-1] {
			break
		}
	}
	root.Left = buildTree(inorder[:index], postorder[:len(inorder[:index])])
	root.Right = buildTree(inorder[index+1:], postorder[len(inorder[:index]):len(inorder[:index])+len(inorder[index+1:])])
	return root
}
