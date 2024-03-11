package leetcode

import "strings"

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.5 MB,击败了10.53% 的Go用户
func capitalizeTitle(title string) string {
	ss := strings.Split(title, " ")
	var result []string
	for index := range ss {
		result = append(result, capitalize(ss[index]))
	}
	return strings.Join(result, " ")
}

func capitalize(title string) string {
	var result []byte
	if len(title) <= 2 {
		for index := range title {
			if title[index] >= 'a' {
				result = append(result, title[index])
			} else {
				result = append(result, title[index]+32)
			}
		}
	} else {
		for index := range title {
			if title[index] >= 'a' {
				result = append(result, title[index])
			} else {
				result = append(result, title[index]+32)
			}
		}
		result[0] = result[0] - 32
	}
	return string(result)
}
