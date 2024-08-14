package leetcode

// 执行耗时:134 ms,击败了78.17% 的Go用户
// 内存消耗:18.7 MB,击败了28.38% 的Go用户
func isArraySpecial(nums []int, queries [][]int) []bool {
	var arraySpecialNums = make([]int, len(nums))
	arraySpecialNums[0] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i]%2 != nums[i-1]%2 {
			arraySpecialNums[i] = arraySpecialNums[i-1] + 1
		} else {
			arraySpecialNums[i] = 1
		}
	}
	var result = make([]bool, len(queries))
	for index, query := range queries {
		result[index] = arraySpecialNums[query[1]]-arraySpecialNums[query[0]] == query[1]-query[0]
	}
	return result
}
