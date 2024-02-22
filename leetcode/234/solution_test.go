package leetcode

import "testing"

func TestSwapPairs(t *testing.T) {
	t.Log(isPalindrome(constructListNode([]int{1, 2, 2, 1})))
	t.Log(isPalindrome(constructListNode([]int{1, 2, 3, 2, 1})))
	t.Log(isPalindrome(constructListNode([]int{1, 2, 1})))
	t.Log(isPalindrome(constructListNode([]int{})))
	t.Log(isPalindrome(constructListNode([]int{1})))
	t.Log(isPalindrome(constructListNode([]int{1, 2, 3, 4, 5})))
	t.Log(isPalindrome(constructListNode([]int{1, 2})))
}
