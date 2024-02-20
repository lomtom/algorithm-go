package leetcode

// 输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
// 输出: [3,9,20,null,null,15,7]
// 执行耗时:3 ms,击败了86.55% 的Go用户
// 内存消耗:4 MB,击败了39.64% 的Go用户
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	num := preorder[0]
	index := 0
	for i := range inorder {
		if inorder[i] == num {
			index = i
			break
		}
	}
	root := &TreeNode{Val: num}
	root.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return root
}
