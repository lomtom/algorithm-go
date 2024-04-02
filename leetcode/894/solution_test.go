package leetcode

import "testing"

func Test(t *testing.T) {
	result := allPossibleFBT(3)
	for index := range result {
		printTree(result[index], 0, "")
	}
	result = allPossibleFBT(7)
	for index := range result {
		printTree(result[index], 0, "")
	}
}
