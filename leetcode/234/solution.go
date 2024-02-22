package leetcode

// 执行耗时:129 ms,击败了26.16% 的Go用户
// 内存消耗:11.1 MB,击败了46.24% 的Go用户
func isPalindrome(head *ListNode) bool {
	var nums = make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}
	return true
}

// 执行耗时:120 ms,击败了52.02% 的Go用户
// 内存消耗:12.6 MB,击败了32.33% 的Go用户
func isPalindrome1(head *ListNode) bool {
	var l int
	var node = head
	for node != nil {
		l++
		node = node.Next
	}
	var stack = make([]int, 0)
	for i := 0; i < l/2; i++ {
		stack = append(stack, head.Val)
		head = head.Next
	}
	if l%2 == 1 {
		head = head.Next
	}
	for i := 0; i < l/2; i++ {
		if stack[len(stack)-1] != head.Val {
			return false
		}
		stack = stack[:len(stack)-1]
		head = head.Next
	}
	return true
}
