package leetcode

//执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：2.9 MB, 在所有 Go 提交中击败了21.52%的用户
func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	stack := []int{}
	for i := len(prices) - 1; i >= 0; i-- {
		for len(stack) > 1 && stack[len(stack)-1] > prices[i] {
			stack = stack[:len(stack)-1]
		}
		ans[i] = prices[i] - stack[len(stack)-1]
		stack = append(stack, prices[i])
	}
	return ans
}
