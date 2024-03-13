package leetcode

import "strings"

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.6 MB,击败了35.90% 的Go用户
func maximumOddBinaryNumber(s string) string {
	var subs [2]int
	for index := range s {
		subs[s[index]-'0']++
	}
	subs[1]--
	var results []byte
	for i := 0; i < subs[1]; i++ {
		results = append(results, '1')
	}
	for i := 0; i < subs[0]; i++ {
		results = append(results, '0')
	}
	results = append(results, '1')
	return string(results)
}

func maximumOddBinaryNumber1(s string) string {
	var countOfOne = strings.Count(s, "1")
	return strings.Repeat("1", countOfOne-1) + strings.Repeat("0", len(s)-countOfOne) + "1"
}
