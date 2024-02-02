package offer

import "algorithm/leetcode"

func reversePrint(head *leetcode.ListNode) []int {
	if head == nil {
		return []int{}
	}
	return append(reversePrint(head.Next), head.Val)
}
