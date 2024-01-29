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

//函数 maximumSwap(num):
//    将 num 转换为字符串 s
//    初始化栈 stack
//    初始化 left 和 right 为字符串 s 的末尾索引
//
//    对于字符串 s 中的每一位数字，从高位向低位遍历:
//        当栈不为空且栈顶元素小于当前元素:
//            弹出栈顶元素，记为 idx
//            如果 idx 小于 left:
//                更新 left 和 right 为 idx 和当前元素的索引
//            如果 idx 等于 right:
//                更新 right 为当前元素的索引
//
//        如果 right 小于字符串 s 的长度且 s[right] 等于当前元素:
//            更新 right 为当前元素的索引
//
//        将当前元素的索引入栈
//
//    如果 left 小于字符串 s 的长度:
//        交换 s[left] 和 s[right]
//        将 s 转换为整数 ans
//
//        返回 ans
//    否则:
//        返回 num

func TestMaximumSwap(t *testing.T) {
	fmt.Println(maximumSwap(1993))
	fmt.Println(maximumSwap(98368))
	fmt.Println(maximumSwap(9972))
	fmt.Println(maximumSwap(9927))
	fmt.Println(maximumSwap(2736))
	fmt.Println(maximumSwap(0))
}
