package leetcode

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:3.8 MB,击败了12.19% 的Go用户
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var root = &ListNode{Val: -101, Next: head}
	remove(root, n)
	return root.Next
}
func remove(head *ListNode, n int) int {
	if head == nil {
		return 0
	}
	var index = remove(head.Next, n)
	if index == n {
		head.Next = head.Next.Next
	}
	return index + 1
}
