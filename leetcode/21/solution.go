package leetcode

// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 执行耗时:2 ms,击败了36.99% 的Go用户
// 内存消耗:2.3 MB,击败了46.62% 的Go用户
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var root = &ListNode{Val: -101}
	var pre = root
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			pre.Next = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			list2 = list2.Next
		}
		pre = pre.Next
	}
	if list1 != nil {
		pre.Next = list1
	}
	if list2 != nil {
		pre.Next = list2
	}
	return root.Next
}
