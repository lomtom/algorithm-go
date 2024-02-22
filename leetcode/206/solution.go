package leetcode

// 输入：head = [nil,1,2,3,4,5] 输出：[5,4,3,2,1]
// 执行耗时:2 ms,击败了28.73% 的Go用户
// 内存消耗:2.4 MB,击败了99.64% 的Go用户
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
