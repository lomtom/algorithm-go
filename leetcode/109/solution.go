package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:5.6 MB,击败了64.48% 的Go用户
func sortedListToBST(head *ListNode) *TreeNode {
	return buildTree(head, nil)
}

func getMid(left, right *ListNode) *ListNode {
	var slow, fast = left, left
	for fast.Next != right && fast.Next.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMid(left, right)
	node := &TreeNode{Val: mid.Val}
	node.Left = buildTree(left, mid)
	node.Right = buildTree(mid.Next, right)
	return node
}
