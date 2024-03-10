package leetcode

import "fmt"

// 执行耗时:0 ms,击败了100.00% 的Go用户
// 内存消耗:2.1 MB,击败了90.20% 的Go用户
func getHint(secret string, guess string) string {
	nums := [10]int{}
	for index := range guess {
		nums[guess[index]-'0']++
	}
	var bulls int
	var cows int
	for index := range guess {
		if secret[index] == guess[index] {
			nums[secret[index]-'0']--
			bulls++
		}
	}
	for index := range guess {
		if secret[index] != guess[index] && nums[secret[index]-'0'] != 0 {
			nums[secret[index]-'0']--
			cows++
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
