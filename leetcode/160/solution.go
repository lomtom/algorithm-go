package leetcode

// 执行耗时:29 ms,击败了45.10% 的Go用户
// 内存消耗:7.1 MB,击败了9.85% 的Go用户
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for headA != nil {
		m[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if m[headB] {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 执行耗时:23 ms,击败了95.66% 的Go用户
// 内存消耗:6.8 MB,击败了91.94% 的Go用户
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
