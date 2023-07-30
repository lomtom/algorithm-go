package leetcode

//执行用时：4 ms, 在所有 Go 提交中击败了100.00%的用户
//内存消耗：5.5 MB, 在所有 Go 提交中击败了8.93%的用户
func canBeEqual(target []int, arr []int) bool {
	m := make(map[int]int, 0)
	for i := 0; i < len(target); i++ {
		m[target[i]]++
		m[arr[i]]--
	}
	for _, value := range m {
		if value != 0 {
			return false
		}
	}
	return true
}
