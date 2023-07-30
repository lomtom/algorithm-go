package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func maximumSwap(num int) (ans int) {
	s := []uint8(strconv.Itoa(num))
	stack := make([]int, 0)
	left, right := len(s)-1, len(s)-1
	for index := range s {
		for len(stack) > 0 && s[stack[len(stack)-1]] < s[index] {
			idx := stack[len(stack)-1]
			if idx < left {
				left, right = idx, index
			}
			if idx == right {
				right = index
			}
			stack = stack[:len(stack)-1]
		}
		if right < len(s) && s[right] == s[index] {
			right = index
		}
		stack = append(stack, index)
	}
	if left < len(s) {
		s[left], s[right] = s[right], s[left]
		ans, _ = strconv.Atoi(string(s))
		return ans
	}

	return num
}

func TestMaximumSwap(t *testing.T) {
	fmt.Println(maximumSwap(1993))
	fmt.Println(maximumSwap(98368))
	fmt.Println(maximumSwap(9972))
	fmt.Println(maximumSwap(9927))
	fmt.Println(maximumSwap(2736))
	fmt.Println(maximumSwap(0))
}
