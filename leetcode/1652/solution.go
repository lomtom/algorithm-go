package leetcode

import (
	"fmt"
	"testing"
)

func decrypt(code []int, k int) []int {
	n := len(code)
	ans := make([]int, n)
	if k == 0 {
		return ans
	}
	code = append(code, code...)
	l, r := 1, k
	if k < 0 {
		l, r = n+k, n-1
	}
	sum := 0
	for _, v := range code[l : r+1] {
		sum += v
	}
	for i := range ans {
		ans[i] = sum
		sum -= code[l]
		sum += code[r+1]
		l, r = l+1, r+1
	}
	return ans
}

func TestDecrypt(t *testing.T) {
	fmt.Println(decrypt([]int{5, 7, 1, 4}, 3))
}
