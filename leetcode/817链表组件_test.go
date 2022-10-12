package leetcode

import (
	"fmt"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) (ans int) {
	m := make(map[int]bool)
	for index := range nums {
		m[nums[index]] = true
	}
	for head != nil {
		if m[head.Val] {
			for head != nil && m[head.Val] {
				head = head.Next
			}
			ans++
		} else {
			head = head.Next
		}
	}
	return
}

func TestNumComponents(t *testing.T) {
	fmt.Println(numComponents(constructListNode([]int{0, 1, 2}), []int{0, 1}))
	fmt.Println(numComponents(constructListNode([]int{0, 1, 2, 3}), []int{0, 1, 3}))
	fmt.Println(numComponents(constructListNode([]int{0, 1, 2, 3, 4, 5}), []int{0, 1, 3, 5}))
	fmt.Println(numComponents(constructListNode([]int{3, 4, 0, 2, 1}), []int{4}))
}
