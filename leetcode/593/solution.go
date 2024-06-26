package leetcode

// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了100.00%的用户
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	if p1[0] == p2[0] && p2[0] == p3[0] && p3[0] == p4[0] && p1[1] == p2[1] && p2[1] == p3[1] && p3[1] == p4[1] {
		return false
	}
	square := func(num int) int {
		return num * num
	}
	q1 := []int{p1[0] - p2[0], p1[1] - p2[1]}
	q2 := []int{p1[0] - p3[0], p1[1] - p3[1]}
	q3 := []int{p1[0] - p4[0], p1[1] - p4[1]}

	s1 := square(q1[0]) + square(q1[1])
	s2 := square(q2[0]) + square(q2[1])
	s3 := square(q3[0]) + square(q3[1])

	c1 := q1[0]*q2[0] + q1[1]*q2[1]
	c2 := q2[0]*q3[0] + q2[1]*q3[1]
	c3 := q1[0]*q3[0] + q1[1]*q3[1]

	if s1 == s2 && s1+s2 == s3 && c1 == 0 && p1[0]-p2[0] == p3[0]-p4[0] && p1[1]-p2[1] == p3[1]-p4[1] ||
		s2 == s3 && s2+s3 == s1 && c2 == 0 && p1[0]-p3[0] == p4[0]-p2[0] && p1[1]-p3[1] == p4[1]-p2[1] ||
		s1 == s3 && s1+s3 == s2 && c3 == 0 && p1[0]-p2[0] == p4[0]-p3[0] && p1[1]-p2[1] == p4[1]-p3[1] {
		return true
	}
	return false
}
