package leetcode

import (
	"fmt"
	"testing"
)

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	nums := make([]int, 0)
	bytes1 := []byte(s1)
	bytes2 := []byte(s2)
	for index := range bytes1 {
		if bytes1[index] == bytes2[index] {
			continue
		}
		nums = append(nums, index)
	}
	return s1 == s2 || len(nums) == 2 && bytes1[nums[0]] == bytes2[nums[1]] && bytes1[nums[1]] == bytes2[nums[0]]
}

func TestAreAlmostEqual(t *testing.T) {
	fmt.Println(areAlmostEqual("bank", "kanb"))
	fmt.Println(areAlmostEqual("kelb", "kelb"))
	// bankbb banbbk
}
