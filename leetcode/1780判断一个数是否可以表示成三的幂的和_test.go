package leetcode

import (
	"github.com/lomtom/go-utils/v2/assert"
	"testing"
)

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.8 MB,击败了47.62% 的Go用户
func checkPowersOfThree(n int) bool {
	mul := 1
	for mul <= n {
		mul *= 3
	}
	mul /= 3
	for mul != 1 || n >= 0 {
		if mul == 0 {
			break
		}
		if n-mul == 0 {
			return true
		} else if n-mul > 0 {
			n -= mul
			mul /= 3
		} else {
			mul /= 3
		}
	}
	return false
}

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:1.8 MB,击败了85.71% 的Go用户
//
//	func checkPowersOfThree(n int) bool {
//		for ; n > 0; n /= 3 {
//			if n%3 == 2 {
//				return false
//			}
//		}
//		return true
//	}

func TestCheckPowersOfThree(t *testing.T) {
	collections := []struct {
		input  int
		output bool
	}{
		{
			27,
			true,
		},
		{
			12,
			true,
		},
		{91,
			true,
		},
		{21,
			false,
		},
	}
	s := assert.NewAssert(t)
	for index := range collections {
		s.Equal(collections[index].output, checkPowersOfThree(collections[index].input))
	}
}
