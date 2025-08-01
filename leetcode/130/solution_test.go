package leetcode

import "testing"

func Test(t *testing.T) {
	nums := [][]byte{
		{'O', 'O'}, {'O', 'O'},
	}
	solve(nums)
	print(t, nums)
	nums = [][]byte{
		{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'},
	}
	solve(nums)
	print(t, nums)
}

func print(t *testing.T, nums [][]byte) {
	for _, v := range nums {
		for _, vv := range v {
			t.Logf("%s", string(vv))
		}
	}
}
