package leetcode

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

// 执行耗时:7 ms,击败了95.24% 的Go用户
// 内存消耗:6.6 MB,击败了95.24% 的Go用户
func getImportance(employees []*Employee, id int) int {
	m := make(map[int]*Employee)
	for _, e := range employees {
		m[e.Id] = e
	}
	return dfs(m, id)
}

func dfs(m map[int]*Employee, id int) int {
	var res = m[id].Importance
	var subordinates = m[id].Subordinates
	for index := range subordinates {
		res += dfs(m, subordinates[index])
	}
	return res
}

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}
