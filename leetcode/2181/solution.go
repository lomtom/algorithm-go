package leetcode

// 执行耗时:365 ms,击败了5.50% 的Go用户
// 内存消耗:30.9 MB,击败了5.51% 的Go用户
func mergeNodes(head *ListNode) *ListNode {
	return merge(head.Next)
}

func merge(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var sum int
	for head.Val != 0 {
		sum += head.Val
		head = head.Next
	}
	if sum > 0 {
		head.Next = &ListNode{
			Val:  sum,
			Next: merge(head.Next),
		}
	}
	return head.Next
}
