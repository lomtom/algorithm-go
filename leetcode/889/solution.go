package leetcode

//- 针对前序遍历后的结果，可以分析出结果的构成为`[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]`
//- 针对中序遍历后的结果，可以分析出结果的构成为`[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]`
//- 针对后序遍历后的结果，可以分析出结果的构成为`[ [左子树的前序遍历结果], [右子树的前序遍历结果], 根节点 ]`
//输入：preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
//输出：[1,2,3,4,5,6,7]

// 执行耗时:4 ms,击败了46.06% 的Go用户
// 内存消耗:2.8 MB,击败了98.18% 的Go用户
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	left := 0
	for i, v := range postorder {
		if v == preorder[1] {
			left = i + 1
		}
	}
	root.Left = constructFromPrePost(preorder[1:left+1], postorder[:left])
	root.Right = constructFromPrePost(preorder[left+1:], postorder[left:len(postorder)-1])
	return root
}
