package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

//执行用时：4 ms, 在所有 Go 提交中击败了8.33%的用户
//内存消耗：1.8 MB, 在所有 Go 提交中击败了91.67%的用户
//func isPrefixOfWord(sentence string, searchWord string) int {
//	ss := strings.Split(sentence, " ")
//	for i, s := range ss {
//		if strings.HasPrefix(s, searchWord) {
//			return i + 1
//		}
//	}
//	return -1
//}

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：1.9 MB, 在所有 Go 提交中击败了12.50%的用户
func isPrefixOfWord(sentence string, searchWord string) int {
	sentence = " " + sentence
	l1 := len(sentence)
	l2 := len(searchWord)
	for i := 1; i < l1-l2+1; i++ {
		if sentence[i-1] == ' ' && sentence[i:i+l2] == searchWord {
			return strings.Count(sentence[:i+l2], " ")
		}
	}
	return -1
}

func TestIsPrefixOfWord(t *testing.T) {
	fmt.Println(isPrefixOfWord("love errichto jonathan dumb", "dumb"))
}
